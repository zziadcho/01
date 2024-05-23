package functions

import (
	"fmt"
	"os"
)

// read font file function
func ReadFontFile(banner string) string {
	file, err := os.ReadFile(banner)
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
	}
	content := string(file[1:])
	return content
}
