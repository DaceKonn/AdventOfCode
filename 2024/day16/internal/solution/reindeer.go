package solution

import (
	"fmt"
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
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h-1, currentNode.w, height, width, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w+1, height, width, facingRight, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w-1, height, width, facingLeft, true)
		case facingRight:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w+1, height, width, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h+1, currentNode.w, height, width, facingDown, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h-1, currentNode.w, height, width, facingUp, true)
		case facingDown:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h+1, currentNode.w, height, width, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w+1, height, width, facingRight, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w-1, height, width, facingLeft, true)
		case facingLeft:
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h, currentNode.w-1, height, width, currentFacing, false)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h+1, currentNode.w, height, width, facingDown, true)
			measurePossibility(currentNode, nodes, visibleNodes, walls, currentNode.h-1, currentNode.w, height, width, facingUp, true)
		}
		lH, lW := getLowest(visibleNodes, nodes)
		if lH == 0 || lW == 0 {
			return
		}
		newNode := &nodes[lH][lW]
		if newNode.h == end.GetH() && newNode.w == end.GetW() {
			nodes[end.GetH()][end.GetW()].visited = true
			score = newNode.weight
			return
		}
		recu(newNode, newNode.sourceFacing)
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
		if nodes[h][w].weight == 0 {
			return
		}
		seatNo++
		walkBackNodes := []node{
			nodes[h+1][w],
			nodes[h][w+1],
			nodes[h-1][w],
			nodes[h][w-1],
		}
		weights := make([]int, 0, 4)
		for _, n := range walkBackNodes {
			if walls[n.h][n.w] {
				continue
			}
			logNode(n, "Near seats")
			weights = append(weights, n.weight)
		}

		sort.Ints(weights)
		log.Debug().Ints("weights", weights).Msg("walkback weights")
		for _, n := range walkBackNodes {
			if walls[n.h][n.w] {
				continue
			}
			if n.weight != weights[0] {
				continue
			}
			walkBack(n.h, n.w)
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

func measurePossibility(currentNode *node, nodes [][]node, visibleNodes map[string]helpers.DefaultPoint, walls [][]bool, rH, rW, height, width, sourceFacing int, isTurn bool) {
	if rH < 0 || rW < 0 || rH >= height || rW >= width {
		log.Debug().Msg("Scan out of bounds")
		return
	}
	if walls[rH][rW] {
		log.Debug().Msg("Scanned wall")
		return
	}
	if nodes[rH][rW].visited {
		logNode(nodes[rH][rW], "Scan already visited")
		return
	}
	potentialWeight := currentNode.weight + 1
	if isTurn {
		potentialWeight += 1000
	}
	if nodes[rH][rW].weight < potentialWeight {
		return
	}
	nodes[rH][rW].weight = potentialWeight
	nodes[rH][rW].sourceH = rH
	nodes[rH][rW].sourceW = rW
	nodes[rH][rW].sourceFacing = sourceFacing
	if _, exists := visibleNodes[nodes[rH][rW].id()]; !exists {
		logNode(nodes[rH][rW], "Scanned new node")
		visibleNodes[nodes[rH][rW].id()] = helpers.NewDefaultPoint(rH, rW)
	}
}

func getLowest(visibleNodes map[string]helpers.DefaultPoint, nodes [][]node) (h, w int) {
	log.Debug().Msg("Getting lowest")
	var currentSelected node
	for _, v := range visibleNodes {
		n := nodes[v.GetH()][v.GetW()]
		if n.visited {
			logNode(n, "Already visited")
			continue
		}
		if currentSelected == (node{}) {
			logNode(n, "Potentially next")
			currentSelected = n
			continue
		}
		if currentSelected.weight < n.weight {
			logNode(n, "Not low enough")
			continue
		}
		currentSelected = n
		logNode(currentSelected, "Newest low")
	}
	logNode(currentSelected, "Selected lowest")
	return currentSelected.h, currentSelected.w
}

func logNode(n node, msg string) {
	log.Debug().
		Int("h", n.h).
		Int("w", n.w).
		Int("weight", n.weight).
		Int("sourceFacing", n.sourceFacing).
		Int("sourceH", n.sourceH).
		Int("sourceW", n.sourceW).
		Bool("visited", n.visited).Msg(msg)
}
