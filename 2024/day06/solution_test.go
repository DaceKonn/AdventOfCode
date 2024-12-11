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

	_, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, floorTiles, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetH() != guard.GetOrigin().GetH()-1 {
			t.Errorf("Guard %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
	for _, ft := range floorTiles {
		if ft.GetOrigin().GetH() == 1 && !ft.GetFlags()["visited"] {
			t.Errorf("Tile %s not marked as visited", ft.GetId())
		}
	}
}

func TestGuardsShouldMoveDown(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'v', 'V', 'v'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}

	_, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, floorTiles, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetH() != guard.GetOrigin().GetH()+1 {
			t.Errorf("Guard %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
	for _, ft := range floorTiles {
		if ft.GetOrigin().GetH() == 1 && !ft.GetFlags()["visited"] {
			t.Errorf("Tile %s not marked as visited", ft.GetId())
		}
	}
}

func TestGuardsShouldMoveRight(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'>', '.', '.'},
		{'>', '.', '.'},
		{'>', '.', '.'},
	}

	_, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, floorTiles, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetW() != guard.GetOrigin().GetW()+1 {
			t.Errorf("Guard %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
	for _, ft := range floorTiles {
		if ft.GetOrigin().GetW() == 1 && !ft.GetFlags()["visited"] {
			t.Errorf("Tile %s not marked as visited", ft.GetId())
		}
	}
}

func TestGuardsShouldMoveLeft(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '.', '<'},
		{'.', '.', '<'},
		{'.', '.', '<'},
	}

	_, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, floorTiles, width, height)
	for _, guard := range guards {
		if guard.GetCurrent().GetW() != guard.GetOrigin().GetW()-1 {
			t.Errorf("Guard %s didn't moved in right direction, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
		}
	}
	for _, ft := range floorTiles {
		if ft.GetOrigin().GetW() == 1 && !ft.GetFlags()["visited"] {
			t.Errorf("Tile %s not marked as visited", ft.GetId())
		}
	}
}

func TestGuardsExitArea(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '^', '.'},
		{'<', '.', '>'},
		{'.', 'v', '.'},
	}

	_, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	moveGuards(guards, floorTiles, width, height)
	for _, guard := range guards {
		if !guard.HasExited() {
			t.Errorf("Guard %s didn't exited the area, origin %s, current %s, exited %v",
				guard.GetId(),
				guard.GetOrigin(),
				guard.GetCurrent(),
				guard.HasExited())
		}
	}
}

func TestGuardsMarkedAsExitShouldntUpdate(t *testing.T) {
	var runeMatrix [][]rune = [][]rune{
		{'.', '^', '.'},
		{'<', '.', '>'},
		{'.', 'v', '.'},
	}

	_, guards, floorTiles := firstLevelScan(runeMatrix, width, height)
	for _, guard := range guards {
		guard.SetExited(true)
	}
	moveGuards(guards, floorTiles, width, height)
	for _, guard := range guards {
		if guard.GetOrigin().GetH() != guard.GetCurrent().GetH() || guard.GetOrigin().GetW() != guard.GetCurrent().GetW() {
			t.Errorf("Guard %s wasn't supposed to move, origin %s, current %s", guard.GetId(), guard.GetOrigin(), guard.GetCurrent())
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
