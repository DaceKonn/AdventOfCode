package main

import (
	"testing"

	"github.com/DaceKonn/AdventOfCode/2024/day08/helpers"
)

func TestShouldNotFindAntenas(t *testing.T) {
	var matrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}
	output := getAntenaGroups(matrix, 3, 3)
	if len(output) > 0 {
		t.Error("Output was expected to be empty")
	}
}

func TestShouldFindAntena(t *testing.T) {
	var matrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', 'A', '.'},
		{'.', '.', '.'},
	}
	output := getAntenaGroups(matrix, 3, 3)
	t.Log(output)
	if len(output) != 1 {
		t.Error("Expected to find 1 antena group")
	}
	v, exists := output['A']
	if !exists {
		t.Fatal("Expected to find antena A")
	}

	t.Log(v)
	if len(v) != 1 {
		t.Fatal("Expected to find 1 antena A")
	}
	if v[0].point.GetH() != 1 || v[0].point.GetW() != 1 {
		t.Fatal("Antena found in wrong location")
	}
}

func TestShouldNotMixAntenas(t *testing.T) {
	var matrix [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', 'A', '.'},
		{'.', 'a', '.'},
	}
	output := getAntenaGroups(matrix, 3, 3)
	t.Log(output)
	if len(output) != 2 {
		t.Error("Expected to find 2 antena groups")
	}
	vUppercase, existsUppercase := output['A']
	vLowercase, existsLowercase := output['a']
	if !existsUppercase || !existsLowercase {
		t.Fatal("Expected to find antena A and/or a")
	}

	t.Log(vUppercase)
	t.Log(vLowercase)
	if len(vUppercase) != 1 || len(vLowercase) != 1 {
		t.Fatal("Expected to find 1 antena A and 1 antena a")
	}
	if vUppercase[0].point.GetH() != 1 || vUppercase[0].point.GetW() != 1 {
		t.Error("Antena A found in wrong location")
	}
	if vLowercase[0].point.GetH() != 2 || vLowercase[0].point.GetW() != 1 {
		t.Error("Antena a found in wrong location")
	}
}

func TestShouldFindTwoAntenasOfOneGroup(t *testing.T) {
	var matrix [][]rune = [][]rune{
		{'.', 'a', '.'},
		{'.', '.', '.'},
		{'.', 'a', '.'},
	}
	output := getAntenaGroups(matrix, 3, 3)
	t.Log(output)
	if len(output) != 1 {
		t.Error("Expected to find 1 antena group")
	}
	vLowercase, existsLowercase := output['a']
	if !existsLowercase {
		t.Fatal("Expected to find antena group a")
	}
	t.Log(vLowercase)
	if len(vLowercase) != 2 {
		t.Fatal("Expected to find 2 antenas a")
	}
	if vLowercase[0].point.GetH() != 0 || vLowercase[0].point.GetW() != 1 {
		t.Error("Antena A found in wrong location")
	}
	if vLowercase[1].point.GetH() != 2 || vLowercase[1].point.GetW() != 1 {
		t.Error("Antena a found in wrong location")
	}
}

func TestShouldMakeSimplePairs(t *testing.T) {
	testData := make(map[rune][]antena)
	tr := 'A'
	testData[tr] = []antena{
		{helpers.NewDefaultPoint(0, 0), tr},
		{helpers.NewDefaultPoint(1, 0), tr},
		{helpers.NewDefaultPoint(0, 1), tr},
		{helpers.NewDefaultPoint(1, 1), tr},
	}

	output := getAntenaPairs(testData)
	t.Log(output)
	if len(output) != 6 {
		t.Fatal("Expected output of len 6")
	}
}

func TestShouldGenerateAntinodes(t *testing.T) {
	testData := make([][]antena, 1, 1)
	tr := 'A'
	testData[0] = []antena{
		{helpers.NewDefaultPoint(3, 4), tr},
		{helpers.NewDefaultPoint(5, 5), tr},
	}
	expected := map[helpers.DefaultPoint]bool{
		helpers.NewDefaultPoint(1, 3): false,
		helpers.NewDefaultPoint(7, 6): false,
	}

	output := generateAntinodes(testData, 10, 10)

	for _, o := range output {
		for k := range expected {
			if o.point.GetH() != k.GetH() && o.point.GetW() != k.GetW() {
				continue
			}
			expected[k] = true
		}
	}

	for k, v := range expected {
		if v {
			t.Log("Antinode\t", k, "\t\tfound")
		} else {
			t.Error("Antinode\t", k, "\t\tnot found")
		}
	}
}

func TestShouldGenerateAntinodesVariantB(t *testing.T) {
	testData := make([][]antena, 1, 1)
	tr := 'A'
	testData[0] = []antena{
		{helpers.NewDefaultPoint(5, 5), tr},
		{helpers.NewDefaultPoint(3, 4), tr},
	}
	expected := map[helpers.DefaultPoint]bool{
		helpers.NewDefaultPoint(1, 3): false,
		helpers.NewDefaultPoint(7, 6): false,
	}

	output := generateAntinodes(testData, 10, 10)

	for _, o := range output {
		for k := range expected {
			if o.point.GetH() != k.GetH() && o.point.GetW() != k.GetW() {
				continue
			}
			expected[k] = true
		}
	}

	for k, v := range expected {
		if v {
			t.Log("Antinode\t", k, "\t\tfound")
		} else {
			t.Error("Antinode\t", k, "\t\tnot found")
		}
	}
}

func TestShouldNotGenerateAntinodesOutOfBounds(t *testing.T) {
	testData := make([][]antena, 1, 1)
	tr := 'A'
	testData[0] = []antena{
		{helpers.NewDefaultPoint(0, 0), tr},
		{helpers.NewDefaultPoint(1, 1), tr},
	}

	output := generateAntinodes(testData, 2, 2)
	if len(output) > 0 {
		t.Error("Expected 0 nodes")
	}
}

func TestShouldGetUnique(t *testing.T) {
	testData := []antinode{
		{make([]antena, 0, 0), helpers.NewDefaultPoint(0, 0)},
		{make([]antena, 0, 0), helpers.NewDefaultPoint(0, 0)},
	}

	output := getUnique(testData)
	if len(output) != 1 {
		t.Error("Unique returned wrong amount of nodes ", output)
	}
}
