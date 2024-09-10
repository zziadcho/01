package main

import (
	// "01/math-skills/common/functions"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Correct Usage: go run . dataFile.txt")
	}
	dataFile := os.Args[1]

	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading  the file", err)
		return
	}

	fmt.Println(data)

}
