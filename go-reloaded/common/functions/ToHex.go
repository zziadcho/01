package reloaded

import (
	"strconv"
	"strings"
)

func ToHex(s string) int {
	decimal, _ := strconv.ParseInt(strings.ToUpper(s), 16, 32)
	return int(decimal)
}
