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

	if !partTwo {
		memory = sortMemory(memory)
	} else {
		memory = sortWithoutFragmenting(memory)
	}
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

func sortWithoutFragmenting(memory []cell) []cell {
	result := make([]cell, len(memory), len(memory))
	copy(result, memory)
	empty := make([]block, 0, 0)
	files := make([]block, 0, 0)

	var transfer func(fileBits, spaceBits []cell)
	transfer = func(fileBits, spaceBits []cell) {
		for indx := len(fileBits) - 1; indx >= 0; indx-- {
			if len(spaceBits) == 0 || spaceBits[0].origin > fileBits[indx].origin {
				break
			}
			result[fileBits[indx].origin] = spaceBits[0]
			result[spaceBits[0].origin] = fileBits[indx]
			spaceBits = spaceBits[1:]
		}
	}

	var currentId int
	var currentBlock block
	var resetBlock func(cCell cell)
	var flushBlock func()

	flushBlock = func() {
		switch currentBlock.state {
		case free:
			empty = append(empty, currentBlock)
		case taken:
			files = append(files, currentBlock)
		}
	}

	resetBlock = func(cCell cell) {
		currentBlock = block{cCell.id, 0, []cell{}, cCell.origin, cCell.state}
	}

	for indx, currentCell := range memory {
		if indx == 0 {
			currentId = currentCell.id
			resetBlock(currentCell)
		}
		if currentId != currentCell.id {
			flushBlock()
			currentId = currentCell.id
			resetBlock(currentCell)
		}
		currentBlock.size++
		currentBlock.cells = append(currentBlock.cells, currentCell)
	}
	flushBlock()
	log.Debug().Str("file blocks", fmt.Sprint(files)).Str("empty blocks", fmt.Sprint(empty)).Msg("Blocks created")

	for indx := len(files) - 1; indx >= 0; indx-- {
		// if len(empty) == 0 || empty[0].origin > files[indx].origin {
		// 	break
		// }
		// result[files[indx].origin] = empty[0]
		// result[empty[0].origin] = files[indx]
		// empty = empty[1:]
		var selectedEmpty int = -1
		for emptyIndx, s := range empty {
			if s.size < files[indx].size {
				continue
			}
			selectedEmpty = emptyIndx
			break
		}
		if selectedEmpty == -1 {
			continue
		}
		transfer(files[indx].cells, empty[selectedEmpty].cells)
		empty[selectedEmpty].size -= files[indx].size
		if empty[selectedEmpty].size == 0 {
			continue
		}
		empty[selectedEmpty].cells = empty[selectedEmpty].cells[files[indx].size:]
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

type block struct {
	id     int
	size   int
	cells  []cell
	origin int
	state  int
}
