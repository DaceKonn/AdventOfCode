package main

import (
	"fmt"
	"strings"
)

const file string = "input_test.txt"

const (
	unknwon = iota
	empty
	symbol
	uncheckedNumber
	partNumber
	floatingNumber
	gear
)

func main() {
	fmt.Println("Advent of Code 2023 - day 03")

	headerLog(0, "read file as matrix")
	matrix, width, height, err := readFileToMatrix(file)
	if err != nil {
		log(err)
		return
	}
	log("width:\t", width)
	log("height:\t", height)
	for _, row := range matrix {
		log(row)
	}

	headerLog(0, "create type matrix")
	headerLog(1, "init empty type matrix")
	typeMatrix := initTypeMatrix(width, height)
	for _, row := range typeMatrix {
		log(row)
	}
	headerLog(1, "first level scan")
	symbolPoints := firstLevelScan(matrix, typeMatrix, width, height)
	for _, row := range typeMatrix {
		log(row)
	}
	log(symbolPoints)

	headerLog(1, "second level scan")
	secondLevelScan(typeMatrix, symbolPoints, width, height)
	for _, row := range typeMatrix {
		log(row)
	}

	headerLog(1, "third level scan")
	thirdLevelScan(typeMatrix, width, height)
	for _, row := range typeMatrix {
		log(row)
	}

	headerLog(0, "read numbers")
	partNumbers := readNumber(matrix, typeMatrix, width, height)
	log("part numbers:\t", partNumbers)

	headerLog(0, "result")
	var result int = 0
	for _, num := range partNumbers {
		result += num
	}
	log(result)
}

type markedRune struct {
	r  rune
	id int
}

func readNumber(matrix [][]rune, typeMatrix [][]int, width, height int) (numbers map[int]int) {
	numbers = make(map[int]int)
	for h := range height {
		rawParts := make([]markedRune, 0)
		for w := range width {
			switch {
			case typeMatrix[h][w] >= 4000:
				rawParts = append(rawParts, markedRune{r: matrix[h][w], id: typeMatrix[h][w] - 4000})
				// default:
				// rawParts = append(rawParts, '.')
			}
		}
		if len(rawParts) == 0 {
			continue
		}
		log(rawParts)
		var currentId int = -1
		currentLine := ""
		for _, mr := range rawParts {
			if currentId == -1 {
				currentId = mr.id
			}
			if mr.id != currentId || h == height-1 {
				log("cl: ", currentLine)
				numbers[currentId] = parseInt(currentLine)
				currentLine = ""
				currentId = mr.id
			}
			currentLine = strings.Join([]string{currentLine, string(mr.r)}, "")
		}
	}
	return numbers
}

func thirdLevelScan(typeMatrix [][]int, width, height int) {
	for h := range height {
		for w := range width {
			if typeMatrix[h][w] != uncheckedNumber {
				continue
			}
			typeMatrix[h][w] = floatingNumber
		}
	}
}

func secondLevelScan(typeMatrix [][]int, symbolPoints []symbolPoint, width, height int) {
	for id, sp := range symbolPoints {
		recaursiveNumberScan(typeMatrix, sp.h-1, sp.w-1, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h-1, sp.w, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h-1, sp.w+1, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h+1, sp.w-1, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h+1, sp.w, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h+1, sp.w+1, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h, sp.w-1, width, height, id)
		recaursiveNumberScan(typeMatrix, sp.h, sp.w+1, width, height, id)
	}
}

func recaursiveNumberScan(typeMatrix [][]int, h, w, width, height, id int) {
	if h < 0 || w < 0 || h >= height || w >= width {
		return
	}
	if typeMatrix[h][w] != uncheckedNumber {
		return
	}
	typeMatrix[h][w] = partNumber*1000 + id
	recaursiveNumberScan(typeMatrix, h, w-1, width, height, id)
	recaursiveNumberScan(typeMatrix, h, w+1, width, height, id)
}

func firstLevelScan(matrix [][]rune, typeMatrix [][]int, width, height int) []symbolPoint {
	symbolPoints := make([]symbolPoint, 0)
	for w := range width {
		for h := range height {
			scanned := unknwon
			switch matrix[h][w] {
			case '.':
				scanned = empty
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				scanned = uncheckedNumber
			case '*':
				scanned = gear
				symbolPoints = append(symbolPoints, symbolPoint{w: w, h: h, symbol: matrix[h][w]})
			default:
				scanned = symbol
			}
			typeMatrix[h][w] = scanned
		}
	}
	return symbolPoints
}

func initTypeMatrix(width, height int) [][]int {
	typeMatrix := make([][]int, height)
	for h := range typeMatrix {
		typeMatrix[h] = make([]int, width)
		for w := range typeMatrix[h] {
			typeMatrix[h][w] = unknwon
		}
	}
	return typeMatrix
}
