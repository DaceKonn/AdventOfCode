package main

import (
	"fmt"

	"github.com/DaceKonn/AdventOfCode/2024/day09/helpers"
	"github.com/rs/zerolog/log"
)

// func runSolution(rawLines []string, partTwo bool) {
func runSolution(runeMatrix [][]rune, width, height int, partTwo bool) {
	memory := getMemory(runeMatrix[0])
	for _, cell := range memory {
		if cell.state == free {
			fmt.Print(".")
			continue
		}
		fmt.Print(cell.id)
	}
	fmt.Println()

	memory = sortMemory(memory)
	var result int = 0
	for indx, cell := range memory {
		if cell.state == free {
			fmt.Print(".")
			continue
		}
		fmt.Print(cell.id)
		result += indx * cell.id
	}
	fmt.Println()
	log.Info().Int("result", result).Msg("Result")
}

func getMemory(line []rune) []cell {
	result := make([]cell, 0, len(line))

	file := true
	id := 0
	indx := 0
	for _, r := range line {
		length := helpers.ParseInt(string(r))
		for range length {
			var newCell cell
			switch file {
			case true:
				newCell = cell{taken, id, indx}
			case false:
				newCell = cell{free, -1, indx}
			}
			result = append(result, newCell)
			indx++
		}
		file = !file
		if !file {
			id++
		}
	}

	return result
}

func sortMemory(memory []cell) []cell {
	result := make([]cell, len(memory), len(memory))
	// var result []cell
	copy(result, memory)
	empty := make([]cell, 0, 0)
	files := make([]cell, 0, 0)

	for _, currentCell := range memory {
		switch currentCell.state {
		case free:
			empty = append(empty, currentCell)
		case taken:
			files = append(files, currentCell)
		}
	}

	for indx := len(files) - 1; indx >= 0; indx-- {
		if len(empty) == 0 || empty[0].origin > files[indx].origin {
			break
		}
		result[files[indx].origin] = empty[0]
		result[empty[0].origin] = files[indx]
		empty = empty[1:]
	}
	return result
}

const (
	free = iota
	taken
)

type cell struct {
	state  int
	id     int
	origin int
}
