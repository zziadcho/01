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
	contentSplit := strings.Split(contentString, " ")
	// contentSplit := strings.Split(contentString, " ")

	/* Text Processing */
	fmt.Println(" ")
	for i, element := range contentSplit {
		switch {

			/////////////// hex ///////////////
		case variables.HexFlag.MatchString(element):
			convertedWord := strconv.Itoa(functions.ToHex(contentSplit[i-1]))
			contentSplit[i-1] = convertedWord

			/////////////// bin ///////////////
		case variables.BinFlag.MatchString(element):
			convertedWord := strconv.Itoa(functions.ToBin(contentSplit[i-1]))
			contentSplit[i-1] = convertedWord

			/////////////// up ///////////////
		case variables.UpFlag.MatchString(element):
			convertedWord := functions.ToUpper(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

			/////////////// low ///////////////
		case variables.LowFlag.MatchString(element):
			convertedWord := functions.ToLower(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

			/////////////// cap ///////////////
		case variables.CapFlag.MatchString(element):
			convertedWord := functions.Capitalize(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

			/////////////// multiple up ///////////////
		case variables.UpFlagMulti.MatchString(element):
			multipiler := contentSplit[i+1]
			multipiler = strings.TrimRight(multipiler, ")")
			multipilerInt, err := strconv.Atoi(multipiler)
			if err != nil {
				fmt.Println("there was a problem converting the Up multiplier to integer:", err)
				continue
			}
			if multipilerInt <= 0 || i-multipilerInt < 0 {
				fmt.Println("the Up multiplier is out of range, lower it to fix the problem")
				continue
			}
			for j := i - multipilerInt; j < i; j++ {
				contentSplit[j] = functions.ToUpper(contentSplit[j])
			}
			convertedWord := functions.ToUpper(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

			/////////////// multiple low ///////////////
		case variables.LowFlagMulti.MatchString(element):
			multipiler := contentSplit[i+1]
			multipiler = strings.TrimRight(multipiler, ")")
			multipilerInt, err := strconv.Atoi(multipiler)
			if err != nil {
				fmt.Println("there was a problem converting the Low multiplier to integer:", err)
				continue
			}
			if multipilerInt <= 0 || i-multipilerInt < 0 {
				fmt.Println("the Low multiplier is out of range, lower it to fix the problem")
				continue
			}
			for j := i - multipilerInt; j < i; j++ {
				contentSplit[j] = functions.ToLower(contentSplit[j])
			}
			convertedWord := functions.ToLower(contentSplit[i-1])
			contentSplit[i-1] = convertedWord

			/////////////// multiple cap ///////////////
		case variables.CapFlagMulti.MatchString(element):
			multipiler := contentSplit[i+1]
			multipiler = strings.TrimRight(multipiler, ")")
			multipilerInt, err := strconv.Atoi(multipiler)
			if err != nil {
				fmt.Println("there was a problem converting the Cap multiplier to integer:", err)
				continue
			}
			if multipilerInt <= 0 || i-multipilerInt < 0 {
				fmt.Println("the Cap multiplier is out of range, lower it to fix the problem")
				continue
			}
			for j := i - multipilerInt; j < i; j++ {
				contentSplit[j] = functions.Capitalize(contentSplit[j])
			}
			convertedWord := functions.Capitalize(contentSplit[i-1])
			contentSplit[i-1] = convertedWord
		}
	}

	/////////////// punctuation ///////////////
	functions.HandlePunctuation(contentSplit)

	/////////////// single quotes ///////////////
	functions.HandleSingleQuote(contentSplit)

	/////////////// definite/indefinite articles ///////////////
	functions.CorrectArticles(contentSplit)

	/* Finalization */
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

	fmt.Println("text has been converted, check result.txt")
}
