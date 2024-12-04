package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"
const dontCom string = "don't()"
const doCom string = "do()"

func main() {
	fmt.Println("Advent of Code 2024 - day 3")

	fmt.Println("\n# Reading file")
	lines, err := readFile(file)
	if err != nil {
		fmt.Println("Failed to read file")
		fmt.Println(err)
	}

	fmt.Println("\n# Build calculation")
	calculation := buildCalculationFromLines(lines)
	fmt.Printf("Calculation: \t%s\n", calculation.print())

	fmt.Println("\n# Result")
	fmt.Println(calculation.calculate())
}

func buildCalculationFromLines(lines []string) (calculation calculator) {
	var base add = newAdd()
	line := strings.Join(lines, "")
	// for _, line := range lines {
	lineCalculation, _ := buildCalculation(line, 0)
	// if !found {
	// fmt.Println("not found in line")
	// continue
	// }
	base.addElement(lineCalculation)
	// }
	return base
}

func buildCalculation(line string, depth int) (calculation calculator, found bool) {
	var base add = newAdd()
	depthLog(fmt.Sprintf("@ building at depth: %d", depth), depth)
	if len(line) == 0 {
		depthLog("line to short", depth)
		return base, false
	}

	nextMulIndex := strings.Index(line, mulCommand())

	nextDont := findNextRecentDont(line, depth)
	if 0 <= nextDont && nextDont < nextMulIndex {
		depthLog("dont before next mul", depth)
		nextDo := findNextRecentDo(line, depth)
		if nextDo == -1 {
			return base, false
		}
		return buildCalculation(line[nextDo:], depth+1)
	}

	if nextMulIndex == -1 {
		depthLog("no more valid mul commands", depth)
		return base, false
	}

	_, after, _ := strings.Cut(line, mulCommand())

	if strings.Index(after, mulOpeners()) != 0 {
		// depthLog("oppener not found", depth)
		// depthLog(fmt.Sprintf("line: %s", after), depth+1)
		// depthLog(fmt.Sprintf("opener: %s", mulOpeners()), depth+1)
		return buildCalculation(after, depth+1)
	}

	after = after[1:]

	leftNum, found := extractStringNumber(after, depth)

	if !found || len(leftNum) > mulMaxNumLength {
		// depthLog("not found LEFT number or number too long", depth)
		return buildCalculation(after, depth+1)
	}

	_, after, _ = strings.Cut(after, leftNum)

	if after[0] != ',' {
		// depthLog("next char not separator", depth)
		return buildCalculation(after, depth+1)
	}

	after = after[1:]

	rightNum, found := extractStringNumber(after, depth)

	if !found || len(rightNum) > mulMaxNumLength {
		// depthLog("not found RIGHT number or number too long", depth)
		buildCalculation(after, depth+1)
	}

	_, after, _ = strings.Cut(after, rightNum)

	if strings.Index(after, mulClosures()) != 0 {
		// depthLog("closer not found", depth)
		return buildCalculation(after, depth+1)
	}

	after = after[1:]

	var multiply multiply = newMultiply()
	multiply.addElement(newElement(parseInt(leftNum)))
	multiply.addElement(newElement(parseInt(rightNum)))

	base.addElement(multiply)

	nextElement, found := buildCalculation(after, depth+1)

	if !found {
		return base, true
	}

	base.addElement(nextElement)
	return base, true
}

func extractStringNumber(fragment string, depth int) (num string, found bool) {
	// depthLog("Extracting number", depth+1)
	// depthLog(fmt.Sprintf("fragment: %s", fragment), depth+2)
	const numbers string = "0123456789"
	var index int = -1
	index = strings.IndexAny(fragment, numbers)
	if index != 0 {
		// depthLog("fragment not starting with number", depth+2)
		return "", false
	}
	num = fragment[0:1]
	nextNum, nextFound := extractStringNumber(fragment[1:], depth)
	if !nextFound {
		return num, true
	}
	num += nextNum
	return num, true
}

func depthLog(message string, depth int) {
	var tab string = strings.Repeat("\t", depth+1)
	fmt.Println(tab, message)
}

func findNextRecentDont(line string, depth int) int {
	return strings.Index(line, dontCom)
	// depthLog("dont - most recent", depth+1)
	// depthLog(fmt.Sprintf("%s", line), depth+1)
	// firstDont := strings.Index(line, dontCom)
	// if firstDont == -1 {
	// 	return -1
	// }
	//
	// firstDo := strings.Index(line[firstDont:], doCom)
	// var mostRecentDont int = -1
	// if firstDo == -1 {
	// 	mostRecentDont = firstDont + strings.LastIndex(line[firstDont:], dontCom)
	// } else {
	// 	firstDo += firstDont
	// 	mostRecentDont = firstDont + strings.LastIndex(line[firstDont:firstDo], dontCom)
	// }
	//
	// depthLog(fmt.Sprintf("first dont: \t%d", firstDont), depth+1)
	// depthLog(fmt.Sprintf("first do: \t%d", firstDo), depth+1)
	// depthLog(fmt.Sprintf("most recent dont: \t%d", mostRecentDont), depth+1)
	// return mostRecentDont
}

func findNextRecentDo(line string, depth int) int {
	return strings.Index(line, doCom)
	// depthLog("do - most recent", depth+1)
	// depthLog(fmt.Sprintf("%s", line), depth+1)
	// firstDo := strings.Index(line, doCom)
	// if firstDo == -1 {
	// 	return -1
	// }
	//
	// firstDont := strings.Index(line[firstDo:], dontCom)
	// var mostRecentDo int = -1
	// if firstDont == -1 {
	// 	mostRecentDo = firstDo + strings.LastIndex(line[firstDo:], doCom)
	// } else {
	// 	firstDont += firstDo
	// 	mostRecentDo = firstDo + strings.LastIndex(line[firstDo:firstDont], doCom)
	// }
	//
	// depthLog(fmt.Sprintf("first do: \t%d", firstDo), depth+1)
	// depthLog(fmt.Sprintf("first dont: \t%d", firstDont), depth+1)
	// depthLog(fmt.Sprintf("most recent do: \t%d", mostRecentDo), depth+1)
	// return mostRecentDo
}
