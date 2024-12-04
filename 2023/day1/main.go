package main

import (
	"fmt"
	"strings"
)

const file string = "input_test_my.txt"
const debugFileContent bool = true

const depth int = 0

var numbersMap map[string]string = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	fmt.Println("Advent of Code 2023 - day 1")

	headerLog(0, "Read file")
	lines, err := readFile(file)
	if err != nil {
		log("!! err reading file")
		log(err)
	}

	debug(debugFileContent, strings.Join(lines, "\n"))

	headerLog(0, "Cleanup lines from wordy numbers")
	lines = cleanupLines(lines, depth+1)

	headerLog(0, "Get numbers in lines")
	numbers := getNumbersFromLines(lines, depth+1)

	headerLog(0, "Result")
	var result int = 0
	for _, num := range numbers {
		result += num
	}
	log(result)
}

func cleanupLines(lines []string, depth int) []string {
	var result []string
	for _, line := range lines {
		newLine := cleanupLineRecursive(line)
		result = append(result, newLine)
		depthLog(depth, fmt.Sprint("old line:", line))
		depthLog(depth, fmt.Sprint("new line:", newLine))
	}
	return result
}

func cleanupLineRecursive(line string) string {
	var newLine string
	lowestIndex := -1
	selectedKey := ""
	for key := range numbersMap {
		firstIndex := strings.Index(newLine, key)
		if firstIndex == -1 || lowestIndex >= 0 && firstIndex > lowestIndex {
			return line
		}
		lowestIndex = firstIndex
		selectedKey = key
	}
	if lowestIndex == -1 {
		return line
	}
	newLine = strings.Replace(newLine, selectedKey, numbersMap[selectedKey], 1)
	return newLine
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
