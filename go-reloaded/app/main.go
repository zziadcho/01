package main

import (
	"01/go-reloaded/common/functions"
	"01/go-reloaded/common/variables"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Checking argument count
	if len(os.Args) != 3 {
		fmt.Println("To run the program use: ./appFile ./inputFile ./outputfile")
		return
	}

	// Assigning file paths
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Opening input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening the input file:", err)
		return
	}
	defer file.Close()

	// Reading input file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the input file:", err)
		return
	}

	// Check if the input file is empty
	if len(content) == 0 {
		fmt.Println("The input file is empty")
		return
	}

	// Converting input file to a string and splitting it into an array
	contentString := string(content)
	contentSplit :=  []string{}
	contentSplit = append(contentSplit, contentString)

	// Text Processing
	for i, element := range contentSplit {
		switch {
		case variables.HexFlag.MatchString(element):
			convertedWord := strconv.Itoa(functions.ToHex(contentSplit[i-1]))
			contentSplit[i-1] = convertedWord

		case variables.BinFlag.MatchString(element):
			convertedWord := strconv.Itoa(functions.ToBin(contentSplit[i-1]))
			contentSplit[i-1] = convertedWord

		case variables.UpFlag.MatchString(element):
			convertedWord := functions.ToUpper(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

		case variables.UpFlagMulti.MatchString(element):
			UpFlagNumberMatch := variables.UpFlagMulti.FindStringSubmatch(element)
			if len(UpFlagNumberMatch) > 1 {
			  UpFlagNumber, _ := strconv.Atoi(UpFlagNumberMatch[1])
			  fmt.Println(UpFlagNumber)
			
			if i-UpFlagNumber >= 0 && UpFlagNumber > 0 {
				for j := i - UpFlagNumber; j < i; j++ {
					contentSplit[j] = functions.ToUpper(contentSplit[j])
				}
			}
		}
		case variables.LowFlag.MatchString(element):
			convertedWord := functions.ToLower(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

		case variables.CapFlag.MatchString(element):
			convertedWord := functions.Capitalize(contentSplit[i-1])
			contentSplit[i-1] = convertedWord
		}
	}

	// Finalization
	contentFinal := []string{}
	for _, element := range contentSplit {
		if element != "" {
			contentFinal = append(contentFinal, element)
		}
	}
	contentRejoin := strings.Join(contentFinal, " ")

	// Creating output file
	outputFileCreate, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating the output file:", err)
		return
	}
	defer outputFileCreate.Close()

	// Writing into output file
	_, err = io.WriteString(outputFileCreate, contentRejoin)
	if err != nil {
		fmt.Println("Error writing into the output file:", err)
		return
	}

	fmt.Print("Processing completed successfully")
}
