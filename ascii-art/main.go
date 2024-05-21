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

		for i, line := range printableSplit {
			if line != "" {
				generatedArt := functions.GenerateAsciiArt(line, fontParse)
				fmt.Println(generatedArt)
			}

			if i < len(printableSplit)-1 {
				fmt.Println()
			}
		}
	} else {
		fmt.Println("something went wrong, check your input")
	}
}
