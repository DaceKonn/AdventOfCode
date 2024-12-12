package main

import (
	"fmt"

	"github.com/DaceKonn/AdventOfCode/2024/day06/helpers"
	// "github.com/rs/zerolog"
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

	log.Info().Msg("Move Guards")
	allowedIterations := 10000
	for activeGuards := len(guards); activeGuards > 0 && allowedIterations > 0; {
		log.Debug().Int("active_guards", activeGuards).Int("remaning-iterations", allowedIterations).Msg("Moving active guards")
		moveGuards(obstacles, guards, floorTiles, width, height, false)
		var newActive int = 0
		for _, guard := range guards {
			if guard.IsDeactivated() {
				continue
			}
			newActive++
		}
		activeGuards = newActive
		allowedIterations--
	}

	log.Info().Msg("Count visited tiles")
	var visitedTiles int = 0
	var vt []helpers.Object
	for _, ft := range floorTiles {
		if !ft.GetFlags()["visited"] {
			continue
		}
		vt = append(vt, ft)
		visitedTiles++
	}
	log.Info().Int("visited-tiles", visitedTiles).Msg("Result")

	log.Info().Msg("Try to find new obstacle")
	var possibleObstacles int = 0
	for indx, ft := range vt {
		// if !ft.GetFlags()["visited"] {
		// 	continue
		// }
		if indx%10 == 0 {
			log.Info().Int("indx", indx).Msg("ft processing")
		}
		tmpFloorTiles := make([]helpers.Object, 0, 0)
		tmpObstacles := make([]helpers.Object, 0, 0)
		tmpGuards := make([]*guard, 0, 0)

		// log.Info().Int("floor index", indx).Int("Possible obstacles", possibleObstacles).Msg("checking possible obstacle placement")

		for _, nft := range floorTiles {
			x := helpers.NewEmptyObject()
			// x.SetFlag("visited", false)
			x.SetFlag("Up", false)
			x.SetFlag("Down", false)
			x.SetFlag("Right", false)
			x.SetFlag("Left", false)
			x.SetOrigin(nft.GetOrigin())
			x.SetCurrent(nft.GetOrigin())
			tmpFloorTiles = append(tmpFloorTiles, x)
		}

		for _, nft := range obstacles {
			x := nft.Copy()
			tmpObstacles = append(tmpObstacles, x)
		}

		for _, nft := range guards {
			x := newEmptyGuard()
			x.SetOrigin(nft.GetOrigin().Copy())
			x.SetCurrent(nft.GetOrigin().Copy())
			x.SetFacing(facingUp)
			tmpGuards = append(tmpGuards, x)
		}

		tmpFloorTiles = remove(tmpFloorTiles, indx)
		tmpObstacle := helpers.NewEmptyObject()
		tmpObstacle.SetId(helpers.NewStringId(ft.GetId().String()))
		tmpObstacle.SetSymbol('#')
		tmpObstacle.SetOrigin(ft.GetOrigin().Copy())
		tmpObstacle.SetCurrent(ft.GetOrigin().Copy())
		tmpObstacles = append(tmpObstacles, tmpObstacle)

		// for _, object := range tmpObstacles {
		// 	helpers.LogObjectInfo(object)
		// }

		allowedIterations = 1000
		// if tmpObstacle.GetId().String() == "6-3" {
		// 	zerolog.SetGlobalLevel(zerolog.DebugLevel)
		// } else {
		// 	zerolog.SetGlobalLevel(zerolog.InfoLevel)
		// }
		for activeGuards := len(tmpGuards); activeGuards > 0 && allowedIterations > 0; {
			log.Debug().Int("active_guards", activeGuards).Int("remaning-iterations", allowedIterations).Msg("Moving active guards")
			moveGuards(tmpObstacles, tmpGuards, tmpFloorTiles, width, height, true)
			var newActive int = 0
			for _, guard := range tmpGuards {
				if guard.GetFlags()["looped"] {
					possibleObstacles++
				}
				if guard.IsDeactivated() {
					// log.Warn().Msg("NOT looped guard")
					continue
				}
				newActive++
			}
			activeGuards = newActive
			allowedIterations--
		}

		if allowedIterations <= 1 {
			// log.Info().Msg("looped guard")
			// helpers.LogObjectInfo(ft)
			// possibleObstacles++
		}
	}

	log.Info().Int("possible-obstacles", possibleObstacles).Msg("Result")
}

func remove(slice []helpers.Object, indx int) []helpers.Object {
	return append(slice[:indx], slice[indx+1:]...)
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

func moveGuards(obstacles []helpers.Object, guards []*guard, floorTiles []helpers.Object, width, height int, loopBreak bool) {
	for _, guard := range guards {
		if guard.IsDeactivated() {
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
			log.Debug().
				Str("id", guard.GetId().String()).
				Msg("Guard left area")
			guard.SetExited(true)
		case false:
			for _, ft := range floorTiles {
				if ft.GetOrigin().GetH() == guard.GetCurrent().GetH() && ft.GetOrigin().GetW() == guard.GetCurrent().GetW() {
					if loopBreak && ft.GetFlags()["visited"] {
						matchFacing := false
						switch guard.GetFacing() {
						case facingUp:
							matchFacing = ft.GetFlags()["Up"] == true
						case facingDown:
							matchFacing = ft.GetFlags()["Down"] == true
						case facingRight:
							matchFacing = ft.GetFlags()["Right"] == true
						case facingLeft:
							matchFacing = ft.GetFlags()["Left"] == true
						}
						if matchFacing {
							guard.SetExited(true)
							guard.GetFlags()["looped"] = true
						}
					}
					ft.SetFlag("visited", true)
					switch guard.GetFacing() {
					case facingUp:
						ft.SetFlag("Up", true)
					case facingDown:
						ft.SetFlag("Down", true)
					case facingRight:
						ft.SetFlag("Right", true)
					case facingLeft:
						ft.SetFlag("Left", true)
					}
				}
			}
		}
	}
}
