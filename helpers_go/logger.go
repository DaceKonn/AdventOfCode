package main

import (
	"fmt"
	"strings"
)

func logAny(a ...any) {
	fmt.Println(a...)
}

func headerLog(depth int, message string) {
	if depth < 0 {
		panic("not expected depth < 0 in header log!")
	}

	var hashes string = strings.Repeat("#", depth+1)
	fmt.Println("\n", hashes, message)
}

func depthLog(depth int, message string) {
	if depth < 0 {
		panic("not expected depth < 0 in depth log!")
	}
	var tabs string = strings.Repeat(" ", depth+1)
	fmt.Println(tabs, message)
}

func debug(active bool, a ...any) {
	if !active {
		return
	}
	logAny(a...)
}
