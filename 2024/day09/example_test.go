package main

import (
	"fmt"
	"testing"
)

func TestShouldSortExample(t *testing.T) {
	testData := []cell{
		{taken, 0, 0},
		{taken, 0, 1},
		{free, -1, 2},
		{free, -1, 3},
		{free, -1, 4},
		{taken, 1, 5},
		{taken, 1, 6},
		{taken, 1, 7},
		{free, -1, 8},
		{free, -1, 9},
		{free, -1, 10},
		{taken, 2, 11},
		{free, -1, 12},
		{free, -1, 13},
		{free, -1, 14},
		{taken, 3, 15},
		{taken, 3, 16},
		{taken, 3, 17},
		{free, -1, 18},
		{taken, 4, 19},
		{taken, 4, 20},
		{free, -1, 21},
		{taken, 5, 22},
		{taken, 5, 23},
		{taken, 5, 24},
		{taken, 5, 25},
		{free, -1, 26},
		{taken, 6, 27},
		{taken, 6, 28},
		{taken, 6, 29},
		{taken, 6, 30},
		{free, -1, 31},
		{taken, 7, 32},
		{taken, 7, 33},
		{taken, 7, 34},
		{free, -1, 35},
		{taken, 8, 36},
		{taken, 8, 37},
		{taken, 8, 38},
		{taken, 8, 39},
		{taken, 9, 40},
		{taken, 9, 41},
	}

	output := sortWithoutFragmenting(testData)
	var tmp string
	for _, o := range output {
		if o.id >= 0 {
			tmp += fmt.Sprint(o.id)
		} else {
			tmp += "."
		}
	}
	t.Log("00...111...2...333.44.5555.6666.777.888899")
	t.Log(tmp)
	t.Log("00992111777.44.333....5555.6666.....8888..")
	if tmp != "00992111777.44.333....5555.6666.....8888.." {
		t.Error("Sorting not matches example")
	}
}
