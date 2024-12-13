package main

import (
	"fmt"
	"os"

	"github.com/DaceKonn/AdventOfCode/2024/day09/helpers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const logStdout bool = true
const debug bool = true
const partTwo bool = false
const file string = "inputs/example.txt"

var logFile *os.File

func main() {
	fmt.Println("Advent of Code - 2024 - day 09")

	logFile = helpers.SetupLogFile()
	defer helpers.CloseLogFile(logFile)
	helpers.ConfigureLogger(logFile, logStdout)

	log.Debug().Str("file", file).Msg("Opening file")

	if !debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	rawLines, err := helpers.ReadFile(file)
	// runeMatrix, width, height, err := helpers.ReadFileToMatrix(file)

	if err != nil {
		log.Fatal().Err(err).Str("file", file).Msg("Failed to read input file")
		return
	}
	log.Debug().Msg("Read file")

	runSolution(rawLines, partTwo)
	// runSolution(runeMatrix,width,height,partTwo)
}
