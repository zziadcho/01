package main

import (
	"01/ascii-art/common/functions"
	"fmt"
	"os"
	"strings"
)

// main block
func main() {
	if len(os.Args) == 3 || len(os.Args) == 2 {
		var fontFile string
		if len(os.Args) == 2 {
			fontFile = functions.ReadFontFile("standard.txt")
		} else {
			if functions.ParseBannerFile(os.Args[2]) {
				fontFile = functions.ReadFontFile(functions.AddTxtExtension(os.Args[2]))
			} else {
				fmt.Println("Usage: go run . [STRING] [BANNER]")
				fmt.Println("EX: go run . something standard")
				os.Exit(1)
			}
		}
		flag := false
		if len(os.Args) == 3 && os.Args[2] == "thinkertoy" {
			flag = true
		}
		fontParse := functions.ParseFont(fontFile,flag)
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
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(1)
	}
}
