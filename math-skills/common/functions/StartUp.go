package functions

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var Numbers []string

func StartUp() {
	if len(os.Args) != 2 {
		fmt.Println("Correct Usage: go run . dataFile.txt")
		os.Exit(1)
	}
	dataFile := os.Args[1]

	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Error opening the file", err)
		os.Exit(1)

	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading  the file", err)
		os.Exit(1)
	}

	Numbers = strings.Fields(string(data))
}
