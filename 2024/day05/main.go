package main

import (
	"fmt"
	"strings"
)

const file string = "input_real.txt"

func main() {
	fmt.Println("Advent of Code 2024 - day 5")

	headerLog(0, "Read raw lines")
	rawLines, err := readFile(file)
	if err != nil {
		log(err)
		return
	}
	logLines(rawLines)

	headerLog(0, "Get rules and updates")
	rules, updates := parseLines(rawLines)

	headerLog(1, "rules result")
	log(rules)

	headerLog(1, "updates result")
	logUpdates(updates)

	headerLog(0, "Evaluate updates")
	evaluateUpdates(updates, rules)

	headerLog(1, "updates evaluation result")
	logUpdates(updates)

	headerLog(0, "Sum middle page numbers")
	sum := sumValidMiddlePages(updates)

	headerLog(1, "Result")
	log(sum)
}

func sumValidMiddlePages(updates []update) int {
	var result int = 0
	for i, update := range updates {
		if !update.valid {
			depthLog(1, fmt.Sprint("update: ", i, "\t\tskip"))
			continue
		}
		depthLog(1, fmt.Sprint("update: ", i, "\t\tsum", "\t\t", update.getMiddlePageNumber()))
		result += update.getMiddlePageNumber()
	}
	return result
}

func evaluateUpdates(updates []update, rules []rule) {
	for updateIndx, update := range updates {
		depthLog(1, fmt.Sprint("evaluating update: ", updateIndx))
		depthLog(2, fmt.Sprint(update.pages))
		for _, rule := range rules {
			depthLog(3, fmt.Sprint("processing rule: ", rule.x, "|", rule.y))
			xValue, xExists := update.indexedPages[rule.x]
			yValue, yExists := update.indexedPages[rule.y]

			depthLog(3, fmt.Sprint("x:", xExists, "|y:", yExists))

			if !xExists || !yExists {
				continue
			}

			depthLog(3, fmt.Sprint("x:", xValue, "|y:", yValue))

			if xValue >= yValue {
				depthLog(4, fmt.Sprint("Update: ", updateIndx, " violation of rule: ", rule.x, "|", rule.y))
				updates[updateIndx].valid = false
			}
		}
	}
}

func parseLines(rawLines []string) (rules []rule, updates []update) {
	var rulesSection bool = true
	for _, line := range rawLines {
		if len(line) == 0 {
			rulesSection = false
			continue
		}
		switch rulesSection {
		case true:
			depthLog(1, "processing rule line")
			split := strings.Split(line, "|")
			x := parseInt(split[0])
			y := parseInt(split[1])
			rules = append(rules, rule{x, y})
		case false:
			depthLog(1, "processing update line")
			update := newEmptyUpdate()
			split := strings.Split(line, ",")
			for i, strNum := range split {
				num := parseInt(strNum)
				update.indexedPages[num] = i
				update.pages = append(update.pages, num)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func logUpdates(updates []update) {
	for i, u := range updates {
		depthLog(1, fmt.Sprint("Update: ", i))
		depthLog(2, fmt.Sprint("valid:", u.valid))
		depthLog(2, fmt.Sprint("pages: ", u.pages))
		depthLog(2, fmt.Sprint("indexed: ", u.indexedPages))
	}
}
