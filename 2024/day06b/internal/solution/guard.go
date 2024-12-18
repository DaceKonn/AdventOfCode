package solution

const (
	facingUp = iota
	facingRight
	facingDown
	facingLeft
)

type guard struct {
	facing int
	h      int
	w      int
}

func (g guard) walkGuard(obstacleMatrix [][]bool, pathMatrix [][]map[int]bool, width, height int) ([][]map[int]bool, bool) {

	looped := false
	var walk func(newH, neawW int)
	walk = func(newH, newW int) {
		if newH < 0 || newW < 0 ||
			newH >= height || newW >= width {
			return
		}
		if obstacleMatrix[newH][newW] {
			g.turn()
		} else {
			g.h = newH
			g.w = newW
		}
		if pathMatrix[g.h][g.w][g.facing] {
			looped = true
			return
		}
		pathMatrix[g.h][g.w][g.facing] = true
		switch g.facing {
		case facingUp:
			walk(g.h-1, g.w)
		case facingRight:
			walk(g.h, g.w+1)
		case facingDown:
			walk(g.h+1, g.w)
		case facingLeft:
			walk(g.h, g.w-1)
		}
	}

	walk(g.h-1, g.w)

	return pathMatrix, looped
}

func (g *guard) turn() {
	switch g.facing {
	case facingUp:
		g.facing = facingRight
	case facingRight:
		g.facing = facingDown
	case facingDown:
		g.facing = facingLeft
	case facingLeft:
		g.facing = facingUp
	}
}
