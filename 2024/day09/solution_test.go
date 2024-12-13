package main

import (
	"testing"
)

func TestShouldParseLine(t *testing.T) {
	line := []rune{'1', '1', '1', '1', '1'}
	var memory []cell = getMemory(line)

	t.Log(memory)
	if len(memory) != len(line) {
		t.Fatal("Should be same length as input line")
	}
	if memory[0].state != taken || memory[2].state != taken || memory[4].state != taken {
		t.Error("Files not in right places")
	}
	if memory[1].state != free || memory[3].state != free {
		t.Error("Free space not in right places")
	}
	if memory[0].id != 0 || memory[2].id != 1 || memory[4].id != 2 {
		t.Error("Files not having right id")
	}
}

func TestShouldParseRightLength(t *testing.T) {
	line := []rune{'2', '3', '1', '0', '1'}

	//00...23
	var memory []cell = getMemory(line)

	t.Log(memory)
	if len(memory) != 7 {
		t.Fatal("Should be of length 6, ", len(memory))
	}
	if memory[0].state != taken || memory[1].state != taken || memory[5].state != taken || memory[6].state != taken {
		t.Error("Files not in right places")
	}
	if memory[2].state != free || memory[3].state != free || memory[4].state != free {
		t.Error("Free space not in right places")
	}
	if memory[0].id != 0 || memory[1].id != 0 || memory[5].id != 1 || memory[6].id != 2 {
		t.Error("Files not having right id")
	}
}

func TestShouldSortMemory(t *testing.T) {
	testData := []cell{
		{taken, 0, 0},
		{free, -1, 1},
		{free, -1, 2},
		{taken, 1, 3},
		{taken, 2, 4},
		{taken, 3, 5},
	}
	output := sortMemory(testData)
	t.Log(output)
	if output[0].id != 0 {
		t.Error("Wrong id at desired place")
	}
	if output[1].id != 3 {
		t.Error("Wrong id at desired place")
	}
	if output[2].id != 2 {
		t.Error("Wrong id at desired place")
	}
	if output[3].id != 1 {
		t.Error("Wrong id at desired place")
	}
	if output[4].id != -1 {
		t.Error("Wrong id at desired place")
	}
	if output[5].id != -1 {
		t.Error("Wrong id at desired place")
	}
}
