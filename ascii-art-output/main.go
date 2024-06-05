package main

import (
	"01/ascii-art/common/functions"
	"fmt"
	"os"
	"strings"
)

func main() {

	var fontFile string
	var printableArgs string
	//var outputFile string
	var fontName string
	
	switch {
	case len(os.Args) == 2:
		fontFile = functions.ReadFontFile("standard.txt")
		printableArgs = os.Args[1]
		fontName = "standard"

	case len(os.Args) == 3 && os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" :
		fontFile = functions.ReadFontFile(functions.AddTxtExtension(os.Args[2]))
		printableArgs = os.Args[1]
		fontName = os.Args[2]

	case len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--output="):
		fontFile = functions.ReadFontFile(functions.AddTxtExtension(os.Args[3]))
		printableArgs = os.Args[2]
		fontName = os.Args[3]
		//outputFile = strings.TrimPrefix(os.Args[1], "--output=")

	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n EX: go run . --output=<fileName.txt> something standard")
	}
	fontParse := functions.ParseFont(fontFile, fontName)
	printableSplit := functions.ArgSplitter(printableArgs)
	if functions.CheckEmpty(printableSplit){
		os.Exit(0)
	} else {
		fmt.Print(functions.GeneratorLoop(printableSplit, fontParse))
	}

}
