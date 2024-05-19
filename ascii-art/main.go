package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadFontFile() string {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
	}
	content := string(file[1:])
	return content
}

func ParseFont(data string) map[rune][]string {
	startChar := ' '
	blocks := strings.Split(data, "\n\n")
	fontMap := make(map[rune][]string)

	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		if len(lines) > 0 {
			char := rune(startChar + rune(i))
			fontMap[char] = lines
		} else {
			fmt.Printf("warning: empty or malformed block at index %d", i)
		}

	}

	return fontMap
}

func GenerateAsciiArt(text string, fontMap map[rune][]string) string {
	var result []string
	
	for i := 0; i <= len(os.Args[1]); i++ {
		
	}

	for i := 0; i <= 8; i++ {
		result = append(result, "")
	}
	
	for _, char := range text {
		if charArt, found := fontMap[char]; found {
			for i, line := range charArt {
				result[i] += line
			}
		}

	}
	return strings.Join(result[:len(result)-1], "\n")
}

func main() {
		fontFile := ReadFontFile()
		fontParse := ParseFont(fontFile)
		
		target := regexp.MustCompile(`\\n`)
		printablesInput := os.Args[1]
		
		lockedIn := target.FindStringIndex(printablesInput)
		
		printables := []string{
			printablesInput[:lockedIn[0]],
			printablesInput[lockedIn[0]:lockedIn[1]],
			printablesInput[lockedIn[1]:],
		}
		
		printablesReady := strings.Join(printables, " ")

		fmt.Println(GenerateAsciiArt(printablesReady, fontParse))
	
}
