package helpers

import (
	"fmt"
	"strings"
)

func LogAny(a ...any) {
	fmt.Println(a...)
}

func HeaderLog(depth int, message string) {
	if depth < 0 {
		panic("not expected depth < 0 in header log!")
	}

	var hashes string = strings.Repeat("#", depth+1)
	fmt.Println("\n", hashes, message)
}

func DepthLog(depth int, message string) {
	if depth < 0 {
		panic("not expected depth < 0 in depth log!")
	}
	var tabs string = strings.Repeat(" ", depth+1)
	fmt.Println(tabs, message)
}

func Debug(active bool, a ...any) {
	if !active {
		return
	}
	LogAny(a...)
}
