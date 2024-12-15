package solution

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func travel(heightMatrix [][]int, startH, startW, height, width int) int {
	score := 0
	visited := make([][]bool, height, height)
	for h := range height {
		visited[h] = make([]bool, width, width)
	}

	var recu func(previousHeight, h, w int)
	recu = func(previousHeight, h, w int) {
		if h < 0 || w < 0 ||
			h >= height || w >= width {
			log.Debug().Msg("walking out of bounds")
			return
		}
		if visited[h][w] {
			log.Debug().Msg("already visited")
			return
		}
		if heightMatrix[h][w] != previousHeight+1 {
			log.Debug().Msg("wrong slope")
			return
		}
		visited[h][w] = true
		if heightMatrix[h][w] != 9 {
			log.Debug().Msg("not a 9")
			recu(heightMatrix[h][w], h-1, w)
			recu(heightMatrix[h][w], h+1, w)
			recu(heightMatrix[h][w], h, w-1)
			recu(heightMatrix[h][w], h, w+1)
			return
		}
		score++
	}
	recu(-1, startH, startW)

	// printVisited(heightMatrix, visited, height, width)

	return score
}

func printVisited(heightMatrix [][]int, visited [][]bool, height, width int) {
	fmt.Println()
	for h := range height {
		for w := range width {
			switch visited[h][w] {
			case true:
				fmt.Print(heightMatrix[h][w])
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
