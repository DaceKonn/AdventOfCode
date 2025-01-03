package solution

import (
	"fmt"
	"math"

	"github.com/DaceKonn/AdventOfCode/2024/day16/helpers"
)

func getMaze(runeMatrix [][]rune, height, width int) (nodes [][]node, walls [][]bool, start, end helpers.DefaultPoint) {
	walls = make([][]bool, height, height)
	nodes = make([][]node, height, height)
	for h := range height {
		walls[h] = make([]bool, width, width)
		nodes[h] = make([]node, width, width)
		for w := range width {
			switch runeMatrix[h][w] {
			case 'S':
				start = helpers.NewDefaultPoint(h, w)
				nodes[h][w] = newNode(h, w)
				nodes[h][w].weight[0] = 0
				nodes[h][w].weight[1] = 0
				nodes[h][w].weight[2] = 0
				nodes[h][w].weight[3] = 0
			case 'E':
				end = helpers.NewDefaultPoint(h, w)
				nodes[h][w] = newNode(h, w)
			case '#':
				walls[h][w] = true
			default:
				nodes[h][w] = newNode(h, w)
			}

		}
	}
	return nodes, walls, start, end
}

type node struct {
	h            int
	w            int
	visited      bool
	sourceH      int
	sourceW      int
	sourceFacing int
	weight       map[int]int
}

func (n node) id() string {
	return fmt.Sprint(n.h, "-", n.w)
}

func newNode(h, w int) node {
	faceWeights := make(map[int]int)
	for i := 0; i < 4; i++ {
		faceWeights[i] = math.MaxInt32
	}
	return node{
		h,
		w,
		false,
		-1,
		-1,
		-1,
		faceWeights,
	}
}
