package main

import (
	"testing"
)

var width, height int = 3, 3

func TestShouldScanFourGuards(t *testing.T) {
	var fiveGuardsRuneMatrix [][]rune = [][]rune{
		{'<', '.', 'v'},
		{'>', '.', '.'},
		{'.', '^', 'V'},
	}

	_, guards, _ := firstLevelScan(fiveGuardsRuneMatrix, width, height)
	if len(guards) != 5 {
		t.Errorf("Didn't detect 4 gaurds, actual: %d", len(guards))
	}
}

func TestGuardShouldFaceUp(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '^', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	if guards[0].facing != facingUp {
		t.Errorf("Guard is not facing up, actual: %s", facingToString(guards[0].facing))
	}
}

func TestGuardShouldFaceDown(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', 'v', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	if guards[0].facing != facingDown {
		t.Errorf("Guard is not facing down, actual: %s", facingToString(guards[0].facing))
	}
}

func TestGuardShouldFaceRight(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '>', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	if guards[0].facing != facingRight {
		t.Errorf("Guard is not facing right, actual: %s", facingToString(guards[0].facing))
	}
}

func TestGuardShouldFaceLeft(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '<', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	if guards[0].facing != facingLeft {
		t.Errorf("Guard is not facing left, actual: %s", facingToString(guards[0].facing))
	}
}
