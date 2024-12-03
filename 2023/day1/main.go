package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"
const debugFileContent bool = true

const depth int = 0

func main() {
	fmt.Println("Advent of Code 2023 - day 1")

	headerLog(0, "Read file")
	lines, err := readFile(file)
	if err != nil {
		log("!! err reading file")
		log(err)
	}

	debug(debugFileContent, strings.Join(lines, "\n"))

	headerLog(0, "Get numbers in lines")
	numbers := getNumbersFromLines(lines, depth+1)

	headerLog(0, "Result")
	var result int = 0
	for _, num := range numbers {
		result += num
	}
	log(result)
}

func getNumbersFromLines(lines []string, depth int) []int {
	const numbers string = "0123456789"
	result := []int{}

	for _, line := range lines {
		firstNumberIndex := strings.IndexAny(line, numbers)
		lastNumberIndex := strings.LastIndexAny(line, numbers)
		if firstNumberIndex == -1 || lastNumberIndex == -1 {
			depthLog(depth, "no number found in line")
		}
		number := parseInt(strings.Join([]string{line[firstNumberIndex : firstNumberIndex+1], line[lastNumberIndex : lastNumberIndex+1]}, ""))
		depthLog(depth, fmt.Sprint("found number:", number))
		result = append(result, number)
	}
	return result
}
