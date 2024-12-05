package main

import (
	"bufio"
	"os"
)

func readFile(fileName string) (lines []string, err error) {
	var file *os.File
	var inErr error
	file, inErr = os.Open(fileName)
	if inErr != nil {
		return lines, inErr
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func readFileToMatrix(fileName string) (runes [][]rune, width, height int, err error) {
	var file *os.File
	var inErr error
	file, inErr = os.Open(fileName)
	if inErr != nil {
		return runes, 0, 0, inErr
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		runes = append(runes, []rune(scanner.Text()))
	}
	return runes, len(runes[0]), len(runes), nil
}
