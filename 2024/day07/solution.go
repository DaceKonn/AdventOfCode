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
		var eq helpers.Multiply = helpers.NewMultiply()
		for _, elInt := range equation.elements {
			eqEl := helpers.NewElement(elInt)
			eqEls = append(eqEls, eqEl)
			eq.AddElement(eqEl)
		}
		log.Debug().
			Int("eq-result", eq.Calculate()).
			Str("equation", eq.Print()).
			Msg("attempt calculation")
	}
}

type equation struct {
	resultString    string
	elementsStrings []string
	result          int
	elements        []int
}
