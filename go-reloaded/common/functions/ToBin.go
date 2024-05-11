package functions

import (
	"strconv"
)

func ToBin(s string) int {
	decimal, _ := strconv.ParseInt(s, 2, 32)
	return int(decimal)
}
