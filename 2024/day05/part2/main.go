package main

import (
	"fmt"
	"sort"
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
	orderRules(rules)
	log(rules)

	headerLog(0, "Get super rules")
	superRules := getSuperRules(rules)

	headerLog(1, "super rules result")
	log(superRules)

	headerLog(1, "updates result")
	logUpdates(updates)

	headerLog(0, "Evaluate updates")
	evaluateUpdates(updates, superRules)

	headerLog(1, "updates evaluation result")
	// logUpdates(updates)

	for i, up := range updates {
		if up.recursionDepth == 0 {
			continue
		}
		log(i, " : ", up.initPages, " : ", up.recursionDepth)
	}

	headerLog(0, "Sum middle page numbers")
	validSum, fixedSum, unfixed := sumMiddlePages(updates)

	headerLog(1, "unfixed")
	log(unfixed)
	headerLog(1, "Valid Result")
	log(validSum)
	headerLog(1, "Fixed Result")
	log(fixedSum)
}

func getSuperRules(rules []rule) map[int][]rule {
	var result map[int][]rule = make(map[int][]rule)
	for _, rule := range rules {
		result[rule.x] = append(result[rule.x], rule)
	}
	return result
}

func sumMiddlePages(updates []update) (validSum, fixedSum, unfixed int) {
	for i, update := range updates {
		if !update.valid && !update.fixed {
			depthLog(1, fmt.Sprint("update: ", i, "\t\tunfixed"))
			unfixed++
			continue
		}
		if !update.valid && update.fixed {
			depthLog(1, fmt.Sprint("update: ", i, "\t\tfixed", "\t\t", update.getMiddlePageNumber()))
			fixedSum += update.getMiddlePageNumber()
			continue
		}
		depthLog(1, fmt.Sprint("update: ", i, "\t\tvalid", "\t\t", update.getMiddlePageNumber()))
		validSum += update.getMiddlePageNumber()
	}
	return validSum, fixedSum, unfixed
}

func evaluateUpdates(updates []update, superRules map[int][]rule) {
	for updateIndx, update := range updates {
		depthLog(1, fmt.Sprint("evaluating update: ", updateIndx))
		depthLog(2, fmt.Sprint(update.pages))
		evaluateUpdateRecursive(updates, updateIndx, superRules, 3, 0)
		depthLog(2, fmt.Sprint(update.pages))
		fmt.Print("\n\n")
	}
}

func evaluateUpdateRecursive(updates []update, updateIndex int, superRules map[int][]rule, logDepth int, recursionDepth int) {
	depthLog(3, fmt.Sprint("recursion depth: ", recursionDepth, " update: ", updateIndex))
	updates[updateIndex].recursionDepth = recursionDepth

	if recursionDepth > 1 {
		depthLog(logDepth, "recursion reached limit")
		return
	}

	for x, rules := range superRules {
		// depthLog(logDepth, fmt.Sprint("processing rule: ", rule.x, "|", rule.y))
		xValue, xExists := updates[updateIndex].indexedPages[x]
		if !xExists {
			continue
		}
		var applicableValues []int
		for _, rule := range rules {
			yValue, yExists := updates[updateIndex].indexedPages[rule.y]
			if !yExists {
				continue
			}
			applicableValues = append(applicableValues, yValue)
			depthLog(logDepth+1, fmt.Sprint("Update: ", updateIndex, " violation of rule: ", rule.x, "|", rule.y))
		}
		// depthLog(logDepth, fmt.Sprint("x:", xExists, "|y:", yExists))

		if len(applicableValues) == 0 {
			continue
		}

		// depthLog(logDepth, fmt.Sprint("x:", xValue, "|y:", yValue))
		sort.Ints(applicableValues)
		yValue := applicableValues[0]

		if xValue >= yValue {
			depthLog(logDepth+2, fmt.Sprint("x:", xValue, "|y:", yValue))
			updates[updateIndex].valid = false
			updates[updateIndex].fixed = false
			updates[updateIndex].fixPageIndex(x, yValue)
			updates[updateIndex].refreshPages()
			continue
		}
		updates[updateIndex].fixed = true
	}

	if !updates[updateIndex].valid && !updates[updateIndex].fixed {
		evaluateUpdateRecursive(updates, updateIndex, superRules, logDepth+1, recursionDepth+1)
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
			// depthLog(1, "processing rule line")
			split := strings.Split(line, "|")
			x := parseInt(split[0])
			y := parseInt(split[1])
			rules = append(rules, rule{x, y})
		case false:
			// depthLog(1, "processing update line")
			update := newEmptyUpdate()
			split := strings.Split(line, ",")
			for i, strNum := range split {
				num := parseInt(strNum)
				update.indexedPages[num] = i
				update.pages = append(update.pages, num)
				update.initPages = append(update.initPages, num)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func logUpdates(updates []update) {
	for i, u := range updates {
		depthLog(1, fmt.Sprint("Update:\t", i))
		depthLog(2, fmt.Sprint("valid:\t", u.valid))
		depthLog(2, fmt.Sprint("fixed:\t", u.fixed))
		depthLog(2, fmt.Sprint("pages:\t", u.pages))
		depthLog(2, fmt.Sprint("indexed:\t", u.indexedPages))
	}
}
