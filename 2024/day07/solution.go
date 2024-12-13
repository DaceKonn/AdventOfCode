package main

import (
	"fmt"
	"strings"

	"github.com/DaceKonn/AdventOfCode/2024/day07/helpers"
	"github.com/rs/zerolog/log"
)

func runSolution(rawLines []string) {
	// func runSolution(runeMatrix [][]rune, width, height int) {
	equations := make([]equation, 0, len(rawLines))
	for _, rLine := range rawLines {
		fisrtLevelSplit := strings.Split(rLine, ": ")
		resultString := fisrtLevelSplit[0]
		elementsStrings := strings.Split(fisrtLevelSplit[1], " ")
		log.Debug().
			Str("result-string", resultString).
			Strs("elements-strings", elementsStrings).
			Msg("Reading line")

		el := make([]int, 0, len(elementsStrings))
		for _, els := range elementsStrings {
			el = append(el, helpers.ParseInt(els))
		}
		equations = append(equations, equation{resultString, elementsStrings, helpers.ParseInt(resultString), el})
	}

	for _, equation := range equations {
		log.Debug().
			Str("equation", fmt.Sprint(equation.resultString, " = ", equation.elementsStrings)).
			Msg("Attempting calculation")

		eqEls := make([]helpers.Element, 0, len(equation.elements))
		for _, elInt := range equation.elements {
			eqEl := helpers.NewElement(elInt)
			eqEls = append(eqEls, eqEl)
		}
	}
}

type equation struct {
	resultString    string
	elementsStrings []string
	result          int
	elements        []int
}

const (
	add = iota
	multiply
)

func calculate(equationElements []helpers.Element, result int, operands []int) int {
	if len(operands) == 0 {
		log.Debug().Msg("Initiating operands")
		for range len(equationElements) - 1 {
			operands = append(operands, add)
		}
	}
	log.Debug().Ints("operands", operands).Msg("Will use operands")

	log.Debug().Msg("calculating")

	var eq helpers.Calculator = equationElements[0]
	allMultiply := true
	for indx, oper := range operands {
		switch oper {
		case add:
			addEq := helpers.NewAdd()
			addEq.AddElement(eq)
			addEq.AddElement(equationElements[indx+1])
			eq = addEq
			allMultiply = false
		case multiply:
			multiplyEq := helpers.NewMultiply()
			multiplyEq.AddElement(eq)
			multiplyEq.AddElement(equationElements[indx+1])
			eq = multiplyEq
		}
	}

	log.Debug().Str("equation", eq.Print()).Msg("Build equation")
	if eq.Calculate() == result {
		return result
	}

	if allMultiply {
		return 0
	}

	for operIndx, oper := range operands {
		if oper == multiply {
			continue
		}
		newOperands := append(operands[:operIndx], multiply)
		newOperands = append(newOperands, operands[operIndx+1:]...)
		recu := calculate(equationElements, result, newOperands)
		if recu == 0 {
			return 0
		}
		return recu
	}

	return 0
}
