package main

import (
	"fmt"
	"math"

	"github.com/DaceKonn/AdventOfCode/2024/day08/helpers"
	"github.com/rs/zerolog/log"
)

// func runSolution(rawLines []string, partTwo bool) {
func runSolution(runeMatrix [][]rune, width, height int, partTwo bool) {
	antenaGroups := getAntenaGroups(runeMatrix, width, height)
	antenaPairs := getAntenaPairs(antenaGroups)
	rawNodes := generateAntinodes(antenaPairs, height, width, partTwo)
	result := getUnique(rawNodes)
	log.Info().Int("result", len(result)).Msg("Result")

	fmt.Println()
	for h := range height {
		for w := range width {
			nodeOut := false
			for _, node := range result {
				if node.point.GetH() != h || node.point.GetW() != w {
					continue
				}
				nodeOut = true
				fmt.Print("#")
				break
			}
			if nodeOut {
				continue
			}
			fmt.Print(string(runeMatrix[h][w]))
		}
		fmt.Println()
	}
}

func getAntenaGroups(runeMatrix [][]rune, width, height int) map[rune][]antena {
	result := make(map[rune][]antena)
	for h := range height {
		for w := range width {
			r := runeMatrix[h][w]
			if r == '.' {
				continue
			}

			_, exists := result[r]
			if !exists {
				result[r] = make([]antena, 0, 1)
			}

			result[r] = append(result[r], antena{helpers.NewDefaultPoint(h, w), r})
		}
	}
	return result
}

func getAntenaPairs(antenas map[rune][]antena) [][]antena {
	result := make([][]antena, 0, 0)
	var permute func(current antena, rest []antena)

	permute = func(current antena, rest []antena) {
		if len(rest) == 0 {
			return
		}
		for _, next := range rest {
			result = append(result, []antena{current, next})
		}
		permute(rest[0], rest[1:])
	}

	for _, v := range antenas {
		permute(v[0], v[1:])
	}

	return result
}

func generateAntinodes(pairs [][]antena, height, width int, partTwo bool) []antinode {
	result := make([]antinode, 0, 0)
	if len(pairs) == 0 {
		panic("not enough pairs")
	}

	for _, pair := range pairs {
		if len(pair) != 2 {
			panic("Pair not a pair")
		}
		x := pair[0]
		y := pair[1]

		dH := int(math.Abs(float64(x.point.GetH() - y.point.GetH())))
		dW := int(math.Abs(float64(x.point.GetW() - y.point.GetW())))

		tmp := getValidAntinodes(x.point.GetH(), x.point.GetW(), y.point.GetH(), y.point.GetW(), dH, dW, height, width, x, y, partTwo)
		if len(tmp) > 0 {
			result = append(result, tmp...)
		}
	}

	return result
}

func getValidAntinodes(xH, xW, yH, yW, dH, dW, height, width int, x, y antena, partTwo bool) []antinode {
	result := make([]antinode, 0, 0)
	tmp := make([]antinode, 0, 0)

	if !partTwo {

		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(xH-dH, xW-dW)})
		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(xH+dH, xW+dW)})
		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(yH-dH, yW-dW)})
		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(yH+dH, yW+dW)})

		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(xH-dH, xW+dW)})
		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(xH+dH, xW-dW)})
		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(yH-dH, yW+dW)})
		tmp = append(tmp, antinode{[]antena{x, y}, helpers.NewDefaultPoint(yH+dH, yW-dW)})

		for _, node := range tmp {
			if node.point.GetH() == xH || node.point.GetH() == yH ||
				node.point.GetW() == xW || node.point.GetW() == yW ||
				node.point.GetH() < 0 || node.point.GetW() < 0 ||
				node.point.GetH() >= height || node.point.GetW() >= width {
				continue
			}
			result = append(result, node)
		}
	} else {
		deltaA := helpers.NewDefaultPoint(xH-yH, xW-yW)
		deltaB := helpers.NewDefaultPoint(0-(xH-yH), 0-(xW-yW))
		log.Debug().Str("deltaA", deltaA.String()).Str("deltaB", deltaB.String()).Msg("Deltas")

		var recur func(delta, point helpers.DefaultPoint)

		recur = func(delta, point helpers.DefaultPoint) {
			log.Debug().Msg("Recursion")
			newH := point.GetH() + delta.GetH()
			newW := point.GetW() + delta.GetW()
			if newH < 0 || newW < 0 ||
				newH >= height || newW >= height {
				return
			}
			newPoint := helpers.NewDefaultPoint(newH, newW)
			tmp = append(tmp, antinode{[]antena{x, y}, newPoint})
			recur(delta, newPoint)
		}

		recur(deltaA, x.point)
		recur(deltaB, y.point)

		result = append(result, tmp...)
		result = append(result, antinode{[]antena{x, y}, x.point})
		result = append(result, antinode{[]antena{x, y}, y.point})
	}

	return result
}

func getUnique(antinodes []antinode) []antinode {
	check := make(map[string]bool)
	result := make([]antinode, 0, 0)
	for _, node := range antinodes {
		id := node.point.String()
		if _, exists := check[id]; exists {
			continue
		}
		result = append(result, node)
		check[id] = true
	}
	return result
}

type antena struct {
	point helpers.DefaultPoint
	r     rune
}

type antinode struct {
	origin []antena
	point  helpers.DefaultPoint
}
