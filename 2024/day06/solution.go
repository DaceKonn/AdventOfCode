package main

import (
	"fmt"

	"github.com/DaceKonn/AdventOfCode/2024/day06/helpers"
	"github.com/rs/zerolog/log"
)

// func runSolution(rawLines []string) {
func runSolution(runeMatrix [][]rune, width, height int) {
	log.Info().Msg("First Level Scan")
	obstacles, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	for _, object := range obstacles {
		helpers.LogObjectDebug(object)
	}
	for _, object := range floorTiles {
		helpers.LogObjectDebug(object)
	}
	for _, object := range guards {
		helpers.LogObjectDebug(object)
	}
}

func firstLevelScan(runeMatrix [][]rune, width, height int) (obstacles []helpers.Object, guards []*guard, floorTiles []helpers.Object) {
	obstacles = make([]helpers.Object, 0, 0)
	floorTiles = make([]helpers.Object, 0, width*height)
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
				tl := helpers.NewEmptyObject()
				tl.SetId(helpers.NewStringId(fmt.Sprint(h, "-", w)))
				tl.SetSymbol('.')
				tl.SetOrigin(helpers.NewDefaultPoint(h, w))
				tl.SetCurrent(helpers.NewDefaultPoint(h, w))
				tl.SetFlag("visited", true)
				floorTiles = append(floorTiles, tl)
			case '#':
				no = helpers.NewEmptyObject()
				obstacles = append(obstacles, no)
			default:
				no = helpers.NewEmptyObject()
				no.SetFlag("visited", false)
				floorTiles = append(floorTiles, no)
			}
			no.SetId(helpers.NewStringId(fmt.Sprint(h, "-", w)))
			no.SetSymbol(symbol)
			no.SetOrigin(helpers.NewDefaultPoint(h, w))
			no.SetCurrent(helpers.NewDefaultPoint(h, w))
		}
	}
	return obstacles, guards, floorTiles
}

func moveGuards(obstacles []helpers.Object, guards []*guard, floorTiles []helpers.Object, width, height int) {
	for _, guard := range guards {
		if guard.HasExited() {
			continue
		}
		newH := guard.GetCurrent().GetH()
		newW := guard.GetCurrent().GetW()
		switch guard.GetFacing() {
		case facingUp:
			newH -= 1
		case facingRight:
			newW += 1
		case facingDown:
			newH += 1
		case facingLeft:
			newW -= 1
		}
		var facingObstacle bool = false
		for _, ob := range obstacles {
			if ob.GetOrigin().GetH() == newH && ob.GetOrigin().GetW() == newW {
				facingObstacle = true
				break
			}
		}

		if facingObstacle {
			var newFacing int = facingUnknown
			switch guard.GetFacing() {
			case facingUp:
				newFacing = facingRight
			case facingRight:
				newFacing = facingDown
			case facingDown:
				newFacing = facingLeft
			case facingLeft:
				newFacing = facingUp
			}
			guard.SetFacing(newFacing)
			continue
		}

		guard.GetCurrent().SetH(newH)
		guard.GetCurrent().SetW(newW)
		guardExited := guard.GetCurrent().GetH() < 0 || guard.GetCurrent().GetW() < 0 || guard.GetCurrent().GetH() == height || guard.GetCurrent().GetW() == width
		switch guardExited {
		case true:
			guard.SetExited(true)
		case false:
			for _, ft := range floorTiles {
				if ft.GetOrigin().GetH() == guard.GetCurrent().GetH() && ft.GetOrigin().GetW() == guard.GetCurrent().GetW() {
					ft.SetFlag("visited", true)
				}
			}
		}
	}
}
