package solution

import "github.com/rs/zerolog/log"

// func RunSolution(rawLines []string, partTwo bool) {
func RunSolution(runeMatrix [][]rune, width, height int, partTwo bool) {
	nodes, walls, start, end := getMaze(runeMatrix, height, width)
	score := solveMaze(nodes, walls, start, end, height, width)
	log.Info().Int("Score", score).Msg("Part 1")
}