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
	validateGuardFacing(guards, facingUp, t)
}

func TestGuardShouldFaceDown(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', 'v', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	validateGuardFacing(guards, facingDown, t)
}

func TestGuardShouldFaceRight(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '>', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	validateGuardFacing(guards, facingRight, t)
}

func TestGuardShouldFaceLeft(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '<', '.'},
	}
	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	validateGuardFacing(guards, facingLeft, t)
}

func TestGuardsShouldMoveUp(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'^', '^', '^'},
	}

	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetH() != guard.GetOrigin().GetH()-1 {
			t.Errorf("Guards %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
}

func TestGuardsShouldMoveDown(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'v', 'V', 'v'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}

	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetH() != guard.GetOrigin().GetH()+1 {
			t.Errorf("Guards %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
}

func TestGuardsShouldMoveRight(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'>', '.', '.'},
		{'>', '.', '.'},
		{'>', '.', '.'},
	}

	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetW() != guard.GetOrigin().GetW()+1 {
			t.Errorf("Guards %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
}

func TestGuardsShouldMoveLeft(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '<'},
		{'.', '.', '<'},
		{'.', '.', '<'},
	}

	_, guards, _ := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetW() != guard.GetOrigin().GetW()-1 {
			t.Errorf("Guards %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
}

func validateGuardFacing(guards []*guard, facing int, t *testing.T) {
	if len(guards) == 0 {
		t.Error("Guard not detected")
		return
	}
	if guards[0].facing != facing {
		t.Errorf("Guard is not %s, actual: %s", facingToString(facing), facingToString(guards[0].facing))
		return
	}
}
