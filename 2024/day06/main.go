package main

import (
	"fmt"
	"os"

	"github.com/DaceKonn/AdventOfCode/2024/day06/helpers"
	"github.com/rs/zerolog/log"
)

const logStdout bool = true
const file string = "inputs/example.txt"

var logFile *os.File

func main() {
	fmt.Println("Advent of Code - 2024 - day 06")

	logFile = helpers.SetupLogFile()
	defer helpers.CloseLogFile(logFile)
	helpers.ConfigureLogger(logFile, logStdout)

	log.Debug().Str("file", file).Msg("Opening file")

	// rawLines, err := helpers.ReadFile(file)
	runeMatrix, width, height, err := helpers.ReadFileToMatrix(file)

	if err != nil {
		log.Fatal().Err(err).Str("file", file).Msg("Failed to read input file")
		return
	}
	log.Debug().Msg("Read file")

	// runSolution(rawLines)
	runSolution(runeMatrix, width, height)
}
