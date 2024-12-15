package solution

import (
	"fmt"

	"github.com/DaceKonn/AdventOfCode/2024/day10/helpers"
)

func getHeightMatrix(runeMatrix [][]rune, width, height int) ([][]int, []helpers.DefaultPoint) {
	matrix := make([][]int, height, height)
	startLocations := make([]helpers.DefaultPoint, 0, 0)
	for h := range height {
		matrix[h] = make([]int, width, width)
		for w := range width {
			v := helpers.ParseInt(string(runeMatrix[h][w]))
			matrix[h][w] = v
			if v == 0 {
				startLocations = append(startLocations, helpers.NewDefaultPoint(h, w))
			}
		}
	}
	return matrix, startLocations
}

func printHeightMatrix(heightMatrix [][]int, width, height int) {
	fmt.Println()
	for h := range height {
		for w := range width {
			fmt.Print(heightMatrix[h][w])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printStartingPoints(locations []helpers.DefaultPoint) {
	for _, dp := range locations {
		fmt.Println(dp.String())
	}
}
