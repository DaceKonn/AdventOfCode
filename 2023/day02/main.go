package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"

const maxRed int = 12
const maxGreen int = 13
const maxBlue int = 14

func main() {
	fmt.Println("Advent of Code 2024 - day 02")

	lines, err := readFile(file)
	if err != nil {
		log(err)
		return
	}
	headerLog(0, "parse lines")
	games := parseLines(lines)
	for _, game := range games {
		log(game)
	}

	headerLog(0, "result")
	var result int = 0
	for _, game := range games {
		if !game.valid {
			continue
		}
		result += game.id
	}
	log(result)
}

func parseLines(lines []string) []game {
	games := make([]game, 0)
	for _, line := range lines {
		newGame := newGame()

		a := strings.Split(line, ":")
		log(a)
		newGame.id = parseInt(strings.Replace(strings.Replace(a[0], "Game ", "", 1), ":", "", 1))

		rawSets := strings.Split(a[1], ";")
		for _, rawSet := range rawSets {
			newSet := newSet()
			rawCubes := strings.Split(rawSet, ", ")
			for _, rawCube := range rawCubes {
				raws := strings.Split(strings.Trim(rawCube, " "), " ")
				switch raws[1] {
				case "green":
					newSet.green = parseInt(raws[0])
					newSet.validGreen = newSet.green <= maxGreen
				case "blue":
					newSet.blue = parseInt(raws[0])
					newSet.validBlue = newSet.blue <= maxBlue
				case "red":
					newSet.red = parseInt(raws[0])
					newSet.validRed = newSet.red <= maxRed
				default:
					log(rawSet)
					log(rawCube)
					log(raws)
					panic("I don't know this cube!")
				}
			}
			newGame.sets = append(newGame.sets, newSet)
			newGame.valid = newGame.valid && newSet.validGreen && newSet.validRed && newSet.validBlue
		}
		games = append(games, newGame)
	}
	return games
}
