package utils

import "strconv"

func StrToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
