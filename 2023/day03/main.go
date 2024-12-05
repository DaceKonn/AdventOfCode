package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"

const (
	unknwon = iota
	empty
	symbol
	uncheckedNumber
	partNumber
	floatingNumber
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
	partNumbers, floatingNumbers := readNumber(matrix, typeMatrix, width, height)
	log("part numbers:\t", partNumbers)
	log("floating numbers:\t", floatingNumbers)

	headerLog(0, "result")
	var result int = 0
	for _, num := range partNumbers {
		result += num
	}
	log(result)
}

func readNumber(matrix [][]rune, typeMatrix [][]int, width, height int) (partNumbers, floatingNumbers []int) {
	for h := range height {
		rawParts := make([]rune, 0)
		rawFloats := make([]rune, 0)
		for w := range width {
			switch typeMatrix[h][w] {
			case partNumber:
				rawParts = append(rawParts, rune(matrix[h][w]))
			case floatingNumber:
				rawFloats = append(rawFloats, rune(matrix[h][w]))
			default:
				rawParts = append(rawParts, '.')
				rawFloats = append(rawFloats, '.')
			}
		}
		partsLine := string(rawParts)
		floatsLine := string(rawFloats)

		log(partsLine)
		log(floatsLine)

		for _, pLine := range strings.Split(partsLine, ".") {
			if len(pLine) == 0 {
				continue
			}
			partNumbers = append(partNumbers, parseInt(pLine))
		}

		for _, fLine := range strings.Split(floatsLine, ".") {
			if len(fLine) == 0 {
				continue
			}
			floatingNumbers = append(floatingNumbers, parseInt(fLine))
		}
	}
	return partNumbers, floatingNumbers
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
	for _, sp := range symbolPoints {
		recaursiveNumberScan(typeMatrix, sp.h-1, sp.w-1, width, height)
		recaursiveNumberScan(typeMatrix, sp.h-1, sp.w, width, height)
		recaursiveNumberScan(typeMatrix, sp.h-1, sp.w+1, width, height)
		recaursiveNumberScan(typeMatrix, sp.h+1, sp.w-1, width, height)
		recaursiveNumberScan(typeMatrix, sp.h+1, sp.w, width, height)
		recaursiveNumberScan(typeMatrix, sp.h+1, sp.w+1, width, height)
		recaursiveNumberScan(typeMatrix, sp.h, sp.w-1, width, height)
		recaursiveNumberScan(typeMatrix, sp.h, sp.w+1, width, height)
	}
}

func recaursiveNumberScan(typeMatrix [][]int, h, w, width, height int) {
	if h < 0 || w < 0 || h >= height || w >= width {
		return
	}
	if typeMatrix[h][w] != uncheckedNumber {
		return
	}
	typeMatrix[h][w] = partNumber
	recaursiveNumberScan(typeMatrix, h, w-1, width, height)
	recaursiveNumberScan(typeMatrix, h, w+1, width, height)
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
			default:
				scanned = symbol
				symbolPoints = append(symbolPoints, symbolPoint{w: w, h: h, symbol: matrix[h][w]})
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
