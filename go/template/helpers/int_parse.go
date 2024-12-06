package helpers

import (
	"fmt"
	"strconv"
)

func ParseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return num
}
