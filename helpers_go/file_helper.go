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
