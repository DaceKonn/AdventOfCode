package main

import (
	"fmt"
	"testing"

	"github.com/DaceKonn/AdventOfCode/2024/day07/helpers"
	"github.com/rs/zerolog/log"
)

func TestShouldCalculateSimpleAdd(t *testing.T) {
	defer fmt.Println("")
	eqEl := []helpers.Element{helpers.NewElement(2), helpers.NewElement(3)}
	result := 5
	calculate := calculate(eqEl, 5, []int{0})
	if calculate != result {
		t.Errorf("Expected different result, got %d, expected %d", calculate, result)
	}
}

func TestShouldCalculateSimpleMultiply(t *testing.T) {
	defer fmt.Println("")
	eqEl := []helpers.Element{helpers.NewElement(2), helpers.NewElement(3)}
	result := 6
	calculate := calculate(eqEl, 6, []int{1})
	if calculate != result {
		t.Errorf("Expected different result, got %d, expected %d", calculate, result)
	}
}

func DisabledTestComplex(t *testing.T) {
	testSamples := make(map[int][]helpers.Element)
	//14
	testSamples[5+2+3+4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//17
	testSamples[5*2+3+4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//25
	//!!
	testSamples[(5+2)*3+4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//40
	//!!
	testSamples[(5+2+3)*4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//34
	testSamples[5*2*3+4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//52
	//!!
	testSamples[(5*2+3)*4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//84
	//!!
	testSamples[(5+2)*3*4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	//120
	testSamples[5*2*3*4] = []helpers.Element{helpers.NewElement(5), helpers.NewElement(2), helpers.NewElement(3), helpers.NewElement(4)}
	for k, v := range testSamples {
		log.Info().Int("result", k).Str("values", fmt.Sprint(v)).Msg("")
		result := calculate(v, k, []int{})
		if result != k {
			t.Errorf("Result \tnot match\t got %d\t expected %d", result, k)
		} else {
			t.Logf("Result \tmatch\t\t got %d\t expected %d", result, k)
		}
	}
}

func TestShouldReturnProperAmountOfPossiblePermutations(t *testing.T) {
	x := generatePermutations(4, false)
	if len(x) != 16 {
		t.Error("Expected 16, got ", len(x))
	}
}

func TestShouldReturnProperPermutations(t *testing.T) {
	testSamples := make(map[string]bool)
	testSamples["00"] = false
	testSamples["01"] = false
	testSamples["10"] = false
	testSamples["11"] = false

	// x := buildPossibleOperands(2)
	x := generatePermutations(2, false)
	if len(x) != 4 {
		t.Error("Expected 4, got ", len(x))
		t.FailNow()
	}
	for _, perm := range x {
		tmp := fmt.Sprintf("%d%d", perm[0], perm[1])
		_, exists := testSamples[tmp]
		if !exists {
			t.Error("Value doesn't exist: ", tmp)
		} else {
			testSamples[tmp] = true
		}
	}

	for k, v := range testSamples {
		if v {
			t.Logf("Value\t%s\tmatched", k)
		} else {
			t.Errorf("Value\t%s\tnot matched", k)
		}
	}
}

func TestXYZ(t *testing.T) {
	if 0%2 != 0 {
		t.Error("error")
	}
}
