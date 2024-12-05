package main

type set struct {
	green      int
	red        int
	blue       int
	validGreen bool
	validBlue  bool
	validRed   bool
}

func newSet() set {
	return set{
		0,
		0,
		0,
		true,
		true,
		true,
	}
}

type game struct {
	id    int
	sets  []set
	valid bool
}

func newGame() game {
	return game{
		id:    0,
		sets:  make([]set, 0),
		valid: true,
	}
}

func (g game) getPower() int {
	maxGreen, maxRed, maxBlue := 0, 0, 0
	for _, singleSet := range g.sets {
		if singleSet.green > maxGreen {
			maxGreen = singleSet.green
		}
		if singleSet.blue > maxBlue {
			maxBlue = singleSet.blue
		}
		if singleSet.red > maxRed {
			maxRed = singleSet.red
		}
	}
	return maxBlue * maxGreen * maxRed
}
