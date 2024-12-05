package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"
const debugFileContent bool = true

const depth int = 0

var numbersMap map[string]string = map[string]string{
	"zero":  "zer",
	"one":   "on",
	"two":   "tw",
	"three": "thr",
	"four":  "four",
	"five":  "fiv",
	"six":   "six",
	"seven": "sev",
	"eight": "eigh",
	"nine":  "ni",
}

var shortNumbersMap map[string]string = map[string]string{
	"zer":  "0",
	"on":   "1",
	"tw":   "2",
	"thr":  "3",
	"four": "4",
	"fiv":  "5",
	"six":  "6",
	"sev":  "7",
	"eigh": "8",
	"ni":   "9",
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
		depthLog(depth, fmt.Sprint("cleaning line:", line))
		newLine := cleanupLineRecursive(line, depth+1)
		result = append(result, newLine)
		depthLog(depth, fmt.Sprint("old line:", line))
		depthLog(depth, fmt.Sprint("new line:", newLine))
	}
	return result
}

func cleanupLineRecursive(line string, logDepth int) string {
	var newLine string
	lowestIndex := -1
	selectedKey := ""
	for key, shortKey := range numbersMap {
		firstIndex := strings.Index(line, key)
		depthLog(logDepth, fmt.Sprint("key:\t", key, "\tindex:\t", firstIndex))
		if firstIndex == -1 || lowestIndex >= 0 && firstIndex > lowestIndex {
			continue
		}
		lowestIndex = firstIndex
		selectedKey = shortKey
	}
	depthLog(logDepth, fmt.Sprint("selectedKey:\t", selectedKey, "\tlowestIndex:\t", lowestIndex))
	if lowestIndex == -1 {
		return line
	}
	newLine = strings.Replace(line, selectedKey, shortNumbersMap[selectedKey], 1)
	return cleanupLineRecursive(newLine, logDepth+1)
	// return newLine
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
