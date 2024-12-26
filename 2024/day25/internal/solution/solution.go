package solution

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// func RunSolution(rawLines []string, partTwo bool) {
func RunSolution(runeMatrix [][]rune, width, height int, partTwo bool) {
	// logRunMatrixLengths(runeMatrix, height)
	keys, locks := parseLocksAndKeys(runeMatrix, height, width)
	possiblePairs := matchKeysAndLocks(keys, locks, width)
	log.Info().Int("Possible pairs", possiblePairs).Msg("Result")
}

func matchKeysAndLocks(keys []key, locks []lock, width int) int {
	possiblePairs := 0
	for _, l := range locks {
		for _, k := range keys {
			match := true
			for w := range width {
				match = match && (k[w]+l[w] <= 5)
			}
			if !match {
				continue
			}
			possiblePairs++
		}
	}
	return possiblePairs
}

type key []int

type lock []int

func parseLocksAndKeys(runeMatrix [][]rune, height, width int) ([]key, []lock) {
	isKey := false
	newObject := []int{-1, -1, -1, -1, -1}

	keys := make([]key, 0, 0)
	locks := make([]lock, 0, 0)

	flush := func() {
		log.Debug().Ints("New object", newObject).Bool("Is key", isKey).Msg("Flushing object")
		switch isKey {
		case true:
			keys = append(keys, newObject)
		case false:
			locks = append(locks, newObject)
		}
		newObject = []int{-1, -1, -1, -1, -1}
	}

	for h := range height {
		if len(runeMatrix[h]) == 0 {
			flush()
			continue
		}
		isKey = runeMatrix[h][0] == '.'
		for w := range width {
			if runeMatrix[h][w] == '.' {
				continue
			}
			newObject[w]++
		}
	}
	flush()
	log.Info().Int("No of keys", len(keys)).Int("No of locks", len(locks)).Msg("Parsed keys and locks")
	return keys, locks
}

func logRunMatrixLengths(runeMatrix [][]rune, height int) {
	for h := range height {
		rowLength := len(runeMatrix[h])
		fmt.Println(rowLength)
	}
}
