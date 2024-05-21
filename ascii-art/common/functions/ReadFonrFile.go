package functions 

import (
	"fmt"
	"os"
)

// read font file function
func ReadFontFile() string {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
	}
	content := string(file[1:])
	return content
}
