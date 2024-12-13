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
	calculate := calculate(eqEl, 5, []int{})
	if calculate != result {
		t.Errorf("Expected different result, got %d, expected %d", calculate, result)
	}
}

func TestShouldCalculateSimpleMultiply(t *testing.T) {
	defer fmt.Println("")
	eqEl := []helpers.Element{helpers.NewElement(2), helpers.NewElement(3)}
	result := 6
	calculate := calculate(eqEl, 6, []int{})
	if calculate != result {
		t.Errorf("Expected different result, got %d, expected %d", calculate, result)
	}
}

func TestComplex(t *testing.T) {
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
