package main

import (
	"sort"
)

type rule struct {
	// before
	x int
	//after
	y int
}

func orderRules(rules []rule) {
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].y > rules[j].y || rules[i].y == rules[j].y && rules[i].x > rules[j].x

		// return rules[i].y < rules[j].y || rules[i].y == rules[j].y && rules[i].x < rules[j].x
	})
}
