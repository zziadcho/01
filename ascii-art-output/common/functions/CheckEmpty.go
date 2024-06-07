package functions

import (
	"fmt"
	"os"
)

func CheckEmpty(args []string, flag bool, fileName string) bool {
	new_nums := 0
	var newLines []byte
	for _, str := range args {
		if str == "" {
			if new_nums != 0 {
				newLines = append(newLines, '\n')
			}
			new_nums++
		}
	}

	if new_nums == len(args) {
		if flag {
			os.WriteFile(fileName, newLines, 0600)
			return true
		}

		for new_nums > 1 {
			fmt.Println()
			new_nums--
		}
		return true
	}
	return false

}
