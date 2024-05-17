package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFontFile(fileName string) (string, error) {
	// reading font file
	file, err := os.ReadFile(os.Args[2] + ".txt")
	if err != nil {
		fmt.Printf("error reading the file: %v", err)
		return "", err
	}
	content := string(file)
	return content, err
}

func ParseFontFile(data string) map[rune][]string {
	// splitting file into blocks, each block contains a letter
	blocks := strings.Split(data, "\n\n")
	startChar := ' '

	// splitting each block to 8 lines and store the first line a map
	// declaring the font map
	fontMap := make(map[rune][]string)

	for i, block := range blocks {
		// splitting blocks by lines
		lines := strings.Split(block, "\n")

		// mapping the current character to its ascii art
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

	for i := 0; i < 8; i++ {
		result = append(result, "")
	}

	for _, char := range text {
		if charArt, found := fontMap[char]; found {
			for i, line := range charArt {
				result[i] += line + " "
			}
		} else {
			for i := 0; i < 8; i++ {
				result[i] += "        "
			}
		}
	}
	return strings.Join(result, "\n")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go [text to print] [font file]")
		return
	}

	text := os.Args[1]
	fontFile := os.Args[2]

	fontData, err := ReadFontFile(fontFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fontMap := ParseFontFile(fontData)
	asciiArt := GenerateAsciiArt(text, fontMap)
	fmt.Println(asciiArt)
}
