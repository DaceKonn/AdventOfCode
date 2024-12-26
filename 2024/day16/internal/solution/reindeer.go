package solution

import (
	"fmt"
	"math"
	"sort"

	"github.com/DaceKonn/AdventOfCode/2024/day16/helpers"
	"github.com/rs/zerolog/log"
)

const (
	facingUp = iota
	facingRight
	facingDown
	facingLeft
)

func solveMaze(nodes [][]node, walls [][]bool, start, end helpers.DefaultPoint, height, width int) (score int, seatNo int) {

	visibleNodes := make(map[string]helpers.DefaultPoint)
	seats := make([][]bool, height, height)
	for h := range height {
		seats[h] = make([]bool, width, width)
	}

	// block := 99999999
	score = 0
	seatNo = 0
	var recu func(currentNode *node, currentFacing int)
	recu = func(currentNode *node, currentFacing int) {
		// if block == 0 {
		// return
		// }
		// block--
		logNode(*currentNode, "Visitng node")
		if currentNode.visited {
			return
		}
		nodes[currentNode.h][currentNode.w].visited = true
		switch currentFacing {
		case facingUp:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h-1, currentNode.w, height, width, currentFacing, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w+1, height, width, currentFacing, facingRight, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w-1, height, width, currentFacing, facingLeft, true)
		case facingRight:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w+1, height, width, currentFacing, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h+1, currentNode.w, height, width, currentFacing, facingDown, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h-1, currentNode.w, height, width, currentFacing, facingUp, true)
		case facingDown:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h+1, currentNode.w, height, width, currentFacing, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w+1, height, width, currentFacing, facingRight, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w-1, height, width, currentFacing, facingLeft, true)
		case facingLeft:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w-1, height, width, currentFacing, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h+1, currentNode.w, height, width, currentFacing, facingDown, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h-1, currentNode.w, height, width, currentFacing, facingUp, true)
		}
		lH, lW, lF := getLowest(visibleNodes, nodes)
		if lH == 0 || lW == 0 {
			return
		}
		newNode := &nodes[lH][lW]
		if newNode.h == end.GetH() && newNode.w == end.GetW() {
			nodes[end.GetH()][end.GetW()].visited = true
			score = newNode.weight[currentFacing]
			return
		}
		recu(newNode, lF)
	}
	recu(&nodes[start.GetH()][start.GetW()], facingRight)

	var walkBack func(h, w int)
	walkBack = func(h, w int) {
		logNode(nodes[h][w], "Check seats for node")
		if seats[h][w] || walls[h][w] {
			logNode(nodes[h][w], "Node not visited or already seated")
			return
		}
		seats[h][w] = true
		seatNo++
		if nodes[h][w].weight[0] == 0 {
			return
		}
		weights := make([]int, 0, 4)
		walkBackNodes := []node{
			nodes[h+1][w],
			nodes[h-1][w],
			nodes[h][w+1],
			nodes[h][w-1],
		}
		for _, n := range walkBackNodes {
			for _, v := range n.weight {
				weights = append(weights, v)
			}
		}

		sort.Ints(weights)
		log.Debug().Ints("weights", weights).Msg("walkback weights")

		for _, n := range walkBackNodes {
			for _, v := range n.weight {
				if v != weights[0] {
					continue
				}
				walkBack(n.h, n.w)
			}
		}
	}
	walkBack(end.GetH(), end.GetW())

	fmt.Println()
	for h := range height {
		for w := range width {
			if seats[h][w] {
				fmt.Print("O")
				continue
			}
			if nodes[h][w].visited {
				fmt.Print("/")
				continue
			}
			if _, exists := visibleNodes[fmt.Sprint(h, "-", w)]; exists {
				fmt.Print("Z")
				continue
			}
			if walls[h][w] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}

	return score, seatNo
}

func measurePossibility(
	currentNode *node,
	nodes [][]node,
	visibleNodes map[string]helpers.DefaultPoint,
	walls [][]bool,
	rH, rW, height, width, currentFacing, newFacing int,
	isTurn bool,
) {
	logNode(*currentNode, "Scanning from point")
	log.Debug().Int("new-facing", newFacing).Msg("With facing")
	if rH < 0 || rW < 0 || rH >= height || rW >= width {
		log.Debug().Msg("Scan out of bounds")
		return
	}
	if walls[rH][rW] {
		log.Debug().Msg("Scanned wall")
		return
	}
	potentialWeight := nodes[currentNode.h][currentNode.w].weight[currentNode.sourceFacing] + 1
	log.Debug().Int("potential", potentialWeight).Msg("Potential weight")
	if isTurn {
		potentialWeight += 1000
	}
	if nodes[rH][rW].visited {
		logNode(nodes[rH][rW], "Scan already visited")
		return
	}
	if nodes[rH][rW].weight[currentFacing] < potentialWeight {
		logNode(nodes[rH][rW], "Already scanned with lower weight")
		return
	}
	nodes[rH][rW].weight[currentFacing] = potentialWeight
	nodes[rH][rW].sourceH = rH
	nodes[rH][rW].sourceW = rW
	nodes[rH][rW].sourceFacing = currentFacing
	if _, exists := visibleNodes[nodes[rH][rW].id()]; !exists {
		logNode(nodes[rH][rW], "Scanned new node")
		visibleNodes[nodes[rH][rW].id()] = helpers.NewDefaultPoint(rH, rW)
	}
}

func getLowest(visibleNodes map[string]helpers.DefaultPoint, nodes [][]node) (h, w, f int) {
	log.Debug().Msg("Getting lowest")
	var currentSelected node
	var currentFacing int
	var currentScore int
	init := true
	for _, v := range visibleNodes {
		n := nodes[v.GetH()][v.GetW()]
		if n.visited {
			logNode(n, "Already visited")
			continue
		}
		if init {
			logNode(n, "Potentially next")
			currentSelected = n
			currentScore, currentFacing = lowestByFacing(n)
			init = false
			continue
		}
		newScore, newFacing := lowestByFacing(n)
		if currentScore < newScore {
			logNode(n, "Not low enough")
			continue
		}
		currentSelected = n
		currentFacing = newFacing
		currentScore = newScore
		logNode(currentSelected, "Newest low")
	}
	logNode(currentSelected, "Selected lowest")
	return currentSelected.h, currentSelected.w, currentFacing
}

func lowestByFacing(n node) (score, facing int) {
	lowestV := math.MaxInt32
	lowestK := -1
	for k, v := range n.weight {
		if v >= lowestV {
			continue
		}
		lowestV = v
		lowestK = k
	}
	return lowestV, lowestK
}

func logNode(n node, msg string) {
	w := ""
	for k, v := range n.weight {
		w += fmt.Sprintf("[%d]-[%d] | ", k, v)
	}
	log.Debug().
		Int("h", n.h).
		Int("w", n.w).
		Str("weight", w).
		Int("sourceFacing", n.sourceFacing).
		Int("sourceH", n.sourceH).
		Int("sourceW", n.sourceW).
		Bool("visited", n.visited).Msg(msg)
}
