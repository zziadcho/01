package main

import (
	"01/ascii-art/common/functions"
	"fmt"
	"os"
	"strings"
)

// main block
func main() {
	fontFile := functions.ReadFontFile()
	fontParse := functions.ParseFont(fontFile)

	if len(os.Args) > 1 {
		printableArgs := os.Args[1]
		printableSplit := strings.Split(printableArgs, `\n`)
		for i := 0; i < len(printableSplit); i++ {
			functions.GenerateAsciiArt(printableSplit[i], fontParse)
		}
	} else {
		fmt.Println("something went wrong, check your input")
	}

}
