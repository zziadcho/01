package main

import (
	"01/ascii-art/common/functions"
	"fmt"
	"os"
	"strings"
)

func AddTxt(arg string) string {
	var c []rune
	for _,r := range arg {
		c = append(c,r)
	}
	txt := ".txt"
	for _,r := range txt {
		c = append(c,r)
	}
	return string(c)
}

func ParseBannerFile(arg string) bool {
	if arg == "standard" || arg == "shadow" || arg == "thinkertoy" {
		return true
	}
	return false
}

// main block

func main() {
	if len(os.Args) == 3 || len(os.Args) == 2 {
		var fontFile string
		if len(os.Args) == 2 {
			fontFile = functions.ReadFontFile("standard.txt")
		} else if len(os.Args) == 3 {
			if ParseBannerFile(os.Args[2]) {
				fontFile = functions.ReadFontFile(AddTxt(os.Args[2]))
			} else {
				fmt.Println("The Banner File you gave is not right")
			}
		}
		fontParse := functions.ParseFont(fontFile)
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
