package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"
const depth int = 0

func main() {
	fmt.Println("Advent of Code 2024 - day 4")

	headerLog(0, "Read file")
	orgLines, err := readFile(file)
	if err != nil {
		log(err)
		return
	}
	debug(true, strings.Join(orgLines, "\n"))

	headerLog(0, "Build check lines")

	headerLog(1, "copy regular lines")
	checkLines := orgLines

	headerLog(1, "add column lines")
	columnLines := getColumnLines(orgLines)
	log(strings.Join(columnLines, "\n"))
	checkLines = append(checkLines, columnLines...)

	headerLog(1, "add slash and backslash lines")

	headerLog(2, "add slash lines")
	slashLines := getSlashLines(orgLines, false)
	log(strings.Join(slashLines, "\n"))
	checkLines = append(checkLines, slashLines...)

	headerLog(2, "add backslash lines")
	backslashLines := getSlashLines(orgLines, true)
	log(strings.Join(backslashLines, "\n"))
	checkLines = append(checkLines, backslashLines...)

	headerLog(1, "add reverse all lines")
	checkLines = append(checkLines, reverseLines(checkLines)...)

	headerLog(1, "check lines generation output")
	debug(true, strings.Join(checkLines, "\n"))

	headerLog(0, "Count xmas")
	result := countXmas(checkLines)
	headerLog(1, "Result")
	log(result)
}

func countXmas(checkLines []string) int {
	var result int = 0
	for _, line := range checkLines {
		result += strings.Count(line, "XMAS")
	}
	return result
}

func getColumnLines(lines []string) []string {
	var runes [][]rune
	var width int = 0
	for i, line := range lines {
		if i == 0 {
			width = len(line)
		}
		runes = append(runes, []rune(line))
	}

	log(width, ":", len(lines), " | ", len(runes), ":", len(runes[0]))
	var columnLines []string
	for y := 0; y < width; y++ {
		var columnRunes []rune
		for x := 0; x < len(lines); x++ {
			columnRunes = append(columnRunes, runes[x][y])
		}
		columnLines = append(columnLines, string(columnRunes))
	}
	return columnLines
}

func getSlashLines(orgLines []string, backslash bool) []string {
	height := len(orgLines)
	var pre, post int
	switch backslash {
	case true:
		pre = height - 1
		post = 0
	case false:
		pre = 0
		post = height - 1
	}

	var shiftLines []string
	for _, line := range orgLines {
		if post < 0 || pre < 0 {
			panic("this shouldn't happen! post < 0 or pre < 0")
		}
		var preSpace string = strings.Repeat(".", pre)
		var postSpace string = strings.Repeat(".", post)
		shiftLine := strings.Join([]string{preSpace, line, postSpace}, "")
		shiftLines = append(shiftLines, shiftLine)
		switch backslash {
		case true:
			pre--
			post++
		case false:
			pre++
			post--
		}
	}

	log(strings.Join(shiftLines, "\n"))
	columnLines := getColumnLines(shiftLines)

	var slashLines []string
	for _, columnLine := range columnLines {
		slashLines = append(slashLines, strings.Trim(columnLine, "."))
	}
	return slashLines
}

func reverseLines(checkLines []string) []string {
	var reverseLines []string
	for _, regularLine := range checkLines {
		if len(regularLine) <= 1 {
			continue
		}
		reverseLines = append(reverseLines, reverseString(regularLine))
	}
	return reverseLines
}
