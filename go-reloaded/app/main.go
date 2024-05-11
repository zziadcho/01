package main

import (
	"fmt"
	"io"
	"os"
	"strings"
    "01/go-reloaded/common/variables"
)

func main() {
	// checking argument number
	if len(os.Args) != 3 {
		fmt.Println("To run the program use: ./appFile ./inputFile ./outputfile")
	}

	// assigning file paths
	inputFile := os.Args[1]
	// outputFile := os.Args[2]

	// opening input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("there was an error opening the input file")
		return
	}
	defer file.Close()

	// read input file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("there was an error reading the input file")
		return
	}

	// file is empty
	if len(content) == 0 {
		fmt.Println("the input file is empty")
		return
	}

	// converting input file to a string
	contentString := string(content)

	// split input file to a Split array
	contentSplit := strings.Split(contentString, " ")

	/* Text Processing Started */

	for i := 0; i <= len(contentSplit); i++ {
		// the flag is the first element
		if contentSplit[i] == "(hex)" || contentSplit[i] == "(bin)" || contentSplit[i] == "(up)" || contentSplit[i] == "(low)" || contentSplit[i] == "(cap)" && i == 0 {
			fmt.Println("the flag cannot be the first element")
		}
	}
	// for i, element := range contentSplit{
    fmt.Println(variables.HexFlagMulti)
	// }


}
