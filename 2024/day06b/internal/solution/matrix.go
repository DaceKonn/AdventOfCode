package solution

import "fmt"

func initMatrixes(width, height int) (obstacleMatrix [][]bool, pathMatrix [][]map[int]bool) {
	obstacleMatrix = make([][]bool, height, height)
	// floorMatrix = make([][]bool, height, height)
	pathMatrix = make([][]map[int]bool, height, height)
	for h := range height {
		obstacleMatrix[h] = make([]bool, width, width)
		// floorMatrix[h] = make([]bool, width, width)
		pathMatrix[h] = make([]map[int]bool, width, width)
		for w := range width {
			pathMatrix[h][w] = make(map[int]bool)
			pathMatrix[h][w][facingUp] = false
			pathMatrix[h][w][facingDown] = false
			pathMatrix[h][w][facingLeft] = false
			pathMatrix[h][w][facingRight] = false
		}
	}

	// return obstacleMatrix, floorMatrix, pathMatrix
	return obstacleMatrix, pathMatrix
}

func printBoolMatrix(matrix [][]bool, width, height int, sign rune) {
	fmt.Println()
	for h := range height {
		for w := range width {
			switch matrix[h][w] {
			case true:
				fmt.Print(string(sign))
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func boolParse(b bool) int {
	switch b {
	case true:
		return 1
	default:
		return 0
	}
}

func printPathMatrix(matrix [][]map[int]bool, width, height int) {
	fmt.Println()

	for h := range height {
		for w := range width {
			tmp := boolParse(matrix[h][w][facingUp]) +
				boolParse(matrix[h][w][facingRight]) +
				boolParse(matrix[h][w][facingLeft]) +
				boolParse(matrix[h][w][facingDown])
			switch tmp > 0 {
			case true:
				fmt.Print(tmp)
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countVisits(matrix [][]map[int]bool, width, height int) int {
	result := 0
	for h := range height {
		for w := range width {
			tmp := boolParse(matrix[h][w][facingUp]) +
				boolParse(matrix[h][w][facingRight]) +
				boolParse(matrix[h][w][facingLeft]) +
				boolParse(matrix[h][w][facingDown])
			if tmp == 0 {
				continue
			}
			result++
		}
	}
	return result
}
