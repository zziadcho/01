package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"01/go-reloaded/common/functions"
	"01/go-reloaded/common/variables"
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
		if i > 0 {
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
				if i+1 < len(contentSplit) {
					multiplier := contentSplit[i+1]
					multiplier = strings.TrimSuffix(multiplier, ")")
					multiplierInt, err := strconv.Atoi(multiplier)
					contentSplit[i+1] = ""
					if err != nil {
						fmt.Println("there was a problem converting the Cap multiplier to integer, correct flag syntax is (up, number):\n", err)
						continue
					}
					if multiplierInt <= 0 || i-multiplierInt < 0 {
						fmt.Println("the Up multiplier is out of range, lower it to fix the problem")
						contentSplit[i+1] = ""
						continue
					}
					for j := i - multiplierInt; j < i; j++ {
						contentSplit[j] = functions.ToUpper(contentSplit[j])
					}
					convertedWord := functions.ToUpper(contentSplit[i-1])
					contentSplit[i-1] = convertedWord
				}

				/////////////// multiple low ///////////////
			case variables.LowFlagMulti.MatchString(element):
				if i+1 < len(contentSplit) {
					multiplier := contentSplit[i+1]
					multiplier = strings.TrimSuffix(multiplier, ")")
					multiplierInt, err := strconv.Atoi(multiplier)
					contentSplit[i+1] = ""
					if err != nil {
						fmt.Println("there was a problem converting the Cap multiplier to integer, correct flag syntax is (low, number):\n", err)
						continue
					}
					if multiplierInt <= 0 || i-multiplierInt < 0 {
						fmt.Println("the Low multiplier is out of range, lower it to fix the problem")
						contentSplit[i+1] = ""
						continue
					}
					for j := i - multiplierInt; j < i; j++ {
						contentSplit[j] = functions.ToLower(contentSplit[j])
					}
					convertedWord := functions.ToLower(contentSplit[i-1])
					contentSplit[i-1] = convertedWord
				}

				/////////////// multiple cap ///////////////
			case variables.CapFlagMulti.MatchString(element):
				if i+1 < len(contentSplit) {
					multiplier := contentSplit[i+1]
					multiplier = strings.TrimSuffix(multiplier, ")")
					multiplierInt, err := strconv.Atoi(multiplier)
					contentSplit[i+1] = ""
					if err != nil {
						fmt.Println("there was a problem converting the Cap multiplier to integer, correct flag syntax is (cap, number):\n", err)
						continue
					}
					if multiplierInt <= 0 || i-multiplierInt < 0 {
						fmt.Println("the Cap multiplier is out of range, lower it to fix the problem")
						continue
					}
					for j := i - multiplierInt; j < i; j++ {
						contentSplit[j] = functions.Capitalize(contentSplit[j])
					}
					convertedWord := functions.Capitalize(contentSplit[i-1])
					contentSplit[i-1] = convertedWord
				}
			}
		}
	}

	/////////////// definite/indefinite articles ///////////////
	articleRegulation := functions.AdjustArticles(contentSplit)

	/////////////// regulation of single quotes ///////////////
	singleQuoteRegulation := functions.HandleSingleQuote(articleRegulation)

	/////////////////  regulation of punctuation ///////////////
	punctuationRegulation := functions.HandlePunctuation(singleQuoteRegulation)

	/* Finalization */
	// clean up
	contentRejoin := strings.Join(punctuationRegulation, " ")
	contentProcessed := functions.RemoveFlagSuffixes(contentRejoin)

	// split the processed text into words
	contentFinal := []string{}
	for _, element := range strings.Split(contentProcessed, " ") {
		if element != "" {
			contentFinal = append(contentFinal, element)
		}
	}

	for i, element := range contentFinal {
		if i <= len(contentFinal) {
			if strings.ContainsAny(element, ")"){
				contentFinal[i] = ""
				fmt.Println("the flag got removed even tho it had no effect, try using the correct syntax: (flag, number)\n")
			}
		}
	}
	readyToOutput := strings.Join(contentFinal, " ")

	// creating output file
	outputFileCreate, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating the output file:", err)
		return
	}
	defer outputFileCreate.Close()

	// writing into output file
	_, err = io.WriteString(outputFileCreate, readyToOutput)
	if err != nil {
		fmt.Println("Error writing into the output file:", err)
		return
	}

	fmt.Println("Text has been converted, check result.txt")
	fmt.Println(" ")
}
