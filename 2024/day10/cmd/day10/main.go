package main

import (
	"fmt"
	"os"

	"github.com/DaceKonn/AdventOfCode/2024/day10/helpers"
	"github.com/DaceKonn/AdventOfCode/2024/day10/internal/solution"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const logStdout bool = true
const debug bool = false
const partTwo bool = true
const file string = "../../inputs/real.txt"

var logFile *os.File

func main() {
	fmt.Println("Advent of Code - 2024 - day 06b")

	logFile = helpers.SetupLogFile()
	defer helpers.CloseLogFile(logFile)
	helpers.ConfigureLogger(logFile, logStdout)

	log.Debug().Str("file", file).Msg("Opening file")

	if !debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	// rawLines, err := helpers.ReadFile(file)
	runeMatrix, width, height, err := helpers.ReadFileToMatrix(file)

	if err != nil {
		log.Fatal().Err(err).Str("file", file).Msg("Failed to read input file")
		return
	}
	log.Debug().Msg("Read file")

	// solution.RunSolution(rawLines, partTwo)
	solution.RunSolution(runeMatrix, width, height, partTwo)
}
