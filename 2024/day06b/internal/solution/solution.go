package solution

import "github.com/rs/zerolog/log"

// func RunSolution(rawLines []string, partTwo bool) {
func RunSolution(runeMatrix [][]rune, width, height int, partTwo bool) {
	obstacleMatrix, pathMatrix := initMatrixes(width, height)

	var g guard
	for w := range width {
		for h := range height {
			switch runeMatrix[h][w] {
			case '#':
				obstacleMatrix[h][w] = true
			case '^':
				pathMatrix[h][w][facingUp] = true
				g = guard{facingUp, h, w}
			default:
				continue
			}
		}
	}

	printBoolMatrix(obstacleMatrix, width, height, '#')
	// printBoolMatrix(floorMatrix, width, height, '+')
	printPathMatrix(pathMatrix, width, height)

	pathMatrix, _ = g.walkGuard(obstacleMatrix, pathMatrix, width, height)

	printPathMatrix(pathMatrix, width, height)
	log.Info().Int("result", countVisits(pathMatrix, width, height)).Msg("Part 1")

	if !partTwo {
		return
	}

	loops := 0
	possibleObst, _ := initMatrixes(width, height)
	for h := range height {
		for w := range width {
			if !anyVisit(pathMatrix[h][w]) {
				continue
			}
			newObstacles, newPath := initMatrixes(width, height)
			copyObstacles(newObstacles, obstacleMatrix)
			newObstacles[h][w] = true
			// printBoolMatrix(newObstacles, width, height, 'X')
			_, looped := g.walkGuard(newObstacles, newPath, width, height)
			if !looped {
				continue
			}
			loops++
			possibleObst[h][w] = true
		}
	}
	printBoolMatrix(possibleObst, width, height, 'O')
	log.Info().Int("result", loops).Msg("Part 2")
}
