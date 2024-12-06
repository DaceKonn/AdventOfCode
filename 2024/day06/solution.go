package main

import (
	"github.com/DaceKonn/AdventOfCode/2024/day06/helpers"
	"github.com/rs/zerolog/log"
)

// func runSolution(rawLines []string) {
func runSolution(runeMatrix [][]rune, width, height int) {
	log.Info().Msg("First Level Scan")
	objects := firstLevelScan(runeMatrix, width, height)
	for _, object := range objects {
		log.Debug().
			Type("type", object).
			Str("rune", string(object.GetSymbol())).
			Msg("Scanned object")
	}
}

func firstLevelScan(runeMatrix [][]rune, width, height int) []helpers.Object {
	result := make([]helpers.Object, 0, width*height)
	for w := range width {
		for h := range height {
			var no helpers.Object
			switch runeMatrix[h][w] {
			case '^', '>', '<', 'V':
				no = newEmptyGuard()
			case '#':
				no = helpers.NewEmptyObject()
			default:
				continue
			}
			no.SetSymbol(runeMatrix[h][w])
			result = append(result, no)
		}
	}
	return result
}
