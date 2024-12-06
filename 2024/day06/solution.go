package main

import (
	"fmt"

	"github.com/DaceKonn/AdventOfCode/2024/day06/helpers"
	"github.com/rs/zerolog/log"
)

// func runSolution(rawLines []string) {
func runSolution(runeMatrix [][]rune, width, height int) {
	log.Info().Msg("First Level Scan")
	obstacles, guards := firstLevelScan(runeMatrix, width, height)
	for _, object := range obstacles {
		helpers.LogObjectDebug(object)
	}
	for _, object := range guards {
		helpers.LogObjectDebug(object)
	}
}

func firstLevelScan(runeMatrix [][]rune, width, height int) (obstacles []helpers.Object, guards []*guard) {
	obstacles = make([]helpers.Object, 0, 0)
	guards = make([]*guard, 0, 0)
	for w := range width {
		for h := range height {
			var no helpers.Object
			symbol := runeMatrix[h][w]
			switch symbol {
			case '^', '>', '<', 'V', 'v':
				if symbol == 'V' {
					symbol = 'v'
				}
				guard := newEmptyGuard()
				guards = append(guards, guard)
				no = guard
			case '#':
				no = helpers.NewEmptyObject()
				obstacles = append(obstacles, no)
			default:
				continue
			}
			no.SetSymbol(symbol)
			no.SetOrigin(helpers.NewDefaultPoint(h, w))
			no.SetCurrent(helpers.NewDefaultPoint(h, w))
			no.SetId(helpers.NewStringId(fmt.Sprint(h, "-", w)))
		}
	}
	return obstacles, guards
}
