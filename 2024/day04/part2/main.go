package main

import (
	"fmt"
	"slices"
	"strings"
)

const file string = "input_real.txt"
const depth int = 0

type point struct {
	x int
	y int
}

func main() {
	fmt.Println("Advent of Code 2024 - day 4")

	headerLog(0, "Read file")
	orgLines, err := readFile(file)
	if err != nil {
		log(err)
		return
	}
	debug(true, strings.Join(orgLines, "\n"))

	headerLog(0, "Get matrix")
	runeMatrix := getMatrix(orgLines)
	logMatrix(runeMatrix)

	headerLog(0, "Get all A positions")
	startPositions := getStartPositions(runeMatrix)
	log(startPositions)

	headerLog(0, "Count X-MAS")
	result := countXMas(runeMatrix, startPositions)

	headerLog(1, "Result")
	log(result)
}

func countXMas(runeMatrix [][]rune, startPositions []point) int {
	var result int = 0

	allowed := []string{
		"MMSS",
		"SSMM",
		"SMSM",
		"MSMS",
	}

	for _, start := range startPositions {
		ul := runeMatrix[start.x-1][start.y-1]
		ll := runeMatrix[start.x+1][start.y-1]
		ur := runeMatrix[start.x-1][start.y+1]
		lr := runeMatrix[start.x+1][start.y+1]
		checkString := string([]rune{ul, ur, ll, lr})
		log(checkString)
		if !slices.Contains(allowed, checkString) {
			continue
		}
		result++
	}
	return result
}

func getStartPositions(runeMatrix [][]rune) []point {
	var points []point
	for x, row := range runeMatrix {
		for y, rune := range row {
			if rune != 'A' || x == 0 || y == 0 || x == len(runeMatrix)-1 || y == len(row)-1 {
				continue
			}
			points = append(points, point{x, y})
		}
	}
	return points
}

func getMatrix(orgLines []string) [][]rune {
	var result [][]rune
	for _, line := range orgLines {
		result = append(result, []rune(line))
	}

	return result
}

func logMatrix(matrix [][]rune) {
	for _, row := range matrix {
		log(string(row))
	}
}
