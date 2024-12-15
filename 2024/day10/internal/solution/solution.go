package solution

import "github.com/rs/zerolog/log"

// func RunSolution(rawLines []string, partTwo bool) {
func RunSolution(runeMatrix [][]rune, width, height int, partTwo bool) {
	heightMatrix, startLocations := getHeightMatrix(runeMatrix, width, height)
	printHeightMatrix(heightMatrix, width, height)
	printStartingPoints(startLocations)
	score := 0
	for _, st := range startLocations {
		v := travel(heightMatrix, st.GetH(), st.GetW(), height, width, partTwo)
		score += v
	}

	log.Info().Int("result", score).Msg("Part 1 result")

}
