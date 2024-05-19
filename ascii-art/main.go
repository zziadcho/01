package main

import (
	"fmt"
	"os"
	"strings"
)

// read font file function
func ReadFontFile() string {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
	}
	content := string(file[1:])
	return content
}

// function responsible for parsong the font file
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

// ascii art generator function
func GenerateAsciiArt(text string, fontMap map[rune][]string) string {
	var result []string

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

// main block
func main() {
	fontFile := ReadFontFile()
	fontParse := ParseFont(fontFile)

	if len(os.Args) > 1 {
		printableArgs := os.Args[1]
		printableArgs = strings.ReplaceAll(printableArgs, `\n`, "\n")
		Generate := GenerateAsciiArt(printableArgs, fontParse)
		fmt.Println(Generate)
	} else {
		fmt.Println("something went wrong, check your input")
	}

}
