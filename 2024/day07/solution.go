package main

import (
	"fmt"
	"strings"

	"github.com/DaceKonn/AdventOfCode/2024/day07/helpers"
	"github.com/rs/zerolog/log"
)

func runSolution(rawLines []string, partTwo bool) {
	// func runSolution(runeMatrix [][]rune, width, height int) {
	equations := make([]equation, 0, len(rawLines))
	permutationMap := make(map[int][][]int)
	for _, rLine := range rawLines {
		fisrtLevelSplit := strings.Split(rLine, ": ")
		resultString := fisrtLevelSplit[0]
		elementsStrings := strings.Split(fisrtLevelSplit[1], " ")
		log.Debug().
			Str("result-string", resultString).
			Strs("elements-strings", elementsStrings).
			Msg("Reading line")

		el := make([]int, 0, len(elementsStrings))
		eqEl := make([]helpers.Element, 0, len(elementsStrings))
		for _, els := range elementsStrings {
			intV := helpers.ParseInt(els)
			el = append(el, intV)
			eqEl = append(eqEl, helpers.NewElement(intV))
		}
		equations = append(equations, equation{resultString, elementsStrings, helpers.ParseInt(resultString), el, eqEl})
		if _, exists := permutationMap[len(elementsStrings)-1]; !exists {
			permutationMap[len(elementsStrings)-1] = generatePermutations(len(elementsStrings)-1, partTwo)
		}
	}

	var result int = 0
	for indx, equation := range equations {
		log.Info().Int("nr", indx).Int("len", len(equations)).Msg("Calculating")
		log.Debug().
			Str("equation", fmt.Sprint(equation.resultString, " = ", equation.elementsStrings)).
			Msg("Attempting calculation")
		for _, permutation := range permutationMap[len(equation.eqElements)-1] {
			calResult := calculate(equation.eqElements, equation.result, permutation)
			if calResult == equation.result {
				result += calResult
				break
			}
		}
	}

	log.Info().Int("result", result).Msg("Result")
}

type equation struct {
	resultString    string
	elementsStrings []string
	result          int
	elements        []int
	eqElements      []helpers.Element
}

const (
	add = iota
	multiply
	concentrate
)

// generatePermutations generates all permutations of a slice of 0s and 1s of a given length
func generatePermutations(length int, partTwo bool) [][]int {
	var result [][]int
	var permute func([]int, int)

	permute = func(current []int, pos int) {
		if pos == length {
			// Make a copy of the current slice and add it to the result
			temp := make([]int, length)
			copy(temp, current)
			result = append(result, temp)
			return
		}
		// Set the current position to 0 and recurse
		current[pos] = add
		permute(current, pos+1)
		// Set the current position to 1 and recurse
		current[pos] = multiply
		permute(current, pos+1)

		if partTwo {
			current[pos] = concentrate
			permute(current, pos+1)
		}
	}

	// Initialize the permutation with the given length
	permute(make([]int, length), 0)
	return result
}

func calculate(equationElements []helpers.Element, result int, operands []int) int {
	if len(operands) == 0 {
		log.Debug().Msg("Initiating operands")
		for range len(equationElements) - 1 {
			operands = append(operands, multiply)
		}
	}
	log.Debug().Ints("operands", operands).Msg("Will use operands")

	log.Debug().Msg("calculating")

	var eq helpers.Calculator = equationElements[0]
	for indx, oper := range operands {
		switch oper {
		case add:
			addEq := helpers.NewAdd()
			addEq.AddElement(eq)
			addEq.AddElement(equationElements[indx+1])
			eq = addEq
		case multiply:
			multiplyEq := helpers.NewMultiply()
			multiplyEq.AddElement(eq)
			multiplyEq.AddElement(equationElements[indx+1])
			eq = multiplyEq
		case concentrate:
			concentrateEq := helpers.NewConcentrate()
			concentrateEq.AddElement(eq)
			concentrateEq.AddElement(equationElements[indx+1])
			eq = concentrateEq
		}

	}

	log.Debug().Str("equation", eq.Print()).Msg("Build equation")
	eqResult := eq.Calculate()
	if eqResult == result {
		return result
	}

	return 0
}
