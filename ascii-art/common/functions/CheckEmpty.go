package functions

import (
	"fmt"
)

func CheckEmpty(s []string) bool {
	new_nums := 0
	for _,str := range s {
		if str == "" {
			new_nums++
		}
	}
	if new_nums == len(s) {
		for new_nums > 1 {
			fmt.Println()
			new_nums--
		}
		return true
	}
	return false
}