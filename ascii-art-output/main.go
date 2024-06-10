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
	var outputFile string
	var fontName string
	var outputFlag bool
	errorMesage := "Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard"

	switch {

	case len(os.Args) < 2:
		fmt.Println(errorMesage)

	case len(os.Args) == 2:
		fontFile = functions.ReadFontFile("standard.txt")
		printableArgs = os.Args[1]
		fontName = "standard"

	case len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--output="):
		fontFile = functions.ReadFontFile("standard.txt")
		printableArgs = os.Args[2]
		fontName = "standard"
		outputFile = strings.TrimPrefix(os.Args[1], "--output=")
		outputFlag = true

	case len(os.Args) == 3 && os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy":
		fontFile = functions.ReadFontFile(functions.AddTxtExtension(os.Args[2]))
		printableArgs = os.Args[1]
		fontName = os.Args[2]

	case len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--output=") && (os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy"):
		fontFile = functions.ReadFontFile(functions.AddTxtExtension(os.Args[3]))
		printableArgs = os.Args[2]
		fontName = os.Args[3]
		outputFile = strings.TrimPrefix(os.Args[1], "--output=")
		outputFlag = true

	default:
		fmt.Println(errorMesage)
		os.Exit(1)
	}

	if len(outputFile) == 0 {
		fmt.Println(errorMesage)
		os.Exit(1)
	}

	if !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println(errorMesage)
		os.Exit(1)
	}

	fontParse := functions.ParseFont(fontFile, fontName)
	//sat this is so stupid instead of spliting in this line you just split in another function call , stop this is so fucking stupid hada ra hwa lkhra b3inih good practice wla zbi ra tl3 lya lkhra bhadshi codi wla 9wed
	printableSplit := functions.ArgSplitter(printableArgs)
	generatedArt := functions.GeneratorLoop(printableSplit, fontParse)

	switch {

	case outputFlag && functions.CheckEmpty(printableSplit, outputFlag, outputFile):
		os.Exit(0)

	case outputFlag:
		os.WriteFile(outputFile, []byte(generatedArt), 0600)

	case functions.CheckEmpty(printableSplit, outputFlag, outputFile):
		os.Exit(0)

	default:
		fmt.Print(generatedArt)
	}

}
