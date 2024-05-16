package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFontFile(fileName string) string {
	// reading font file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		return ""
	}
	content := string(file)
	return content
}

func ParseFontFile(data string) map[rune][]string {
	// splitting file into blocks, each block contains a letter
	blocks := strings.Split(data, "\n\n")
	startChar := ' '

	// splitting each block to 8 lines and store the first line a map
	// declaring the font map
	fontMap := make(map[rune][]string)

	for i, block := range blocks {
		// removing white spaces from each block
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
	if len(text) <= 12 {
		for i := 0; i <= 8; i++ {
			result = append(result, "")
		}

		for _, char := range text {
			if charArt, found := fontMap[char]; found {
				for i, line := range charArt {
					result[i] += line + " "
				}
			} else {
				for i := 0; i <= 8; i++ {
					result[i] += "        "
				}
			}
		}
	}else{
    fmt.Printf("string length has to be below 12, your input is: %v", len(text))
  }
  return strings.Join(result, "\n")
}

func main() {
	fontData := ReadFontFile("standard.txt")
	fontMap := ParseFontFile(fontData)
	generatedArt := GenerateAsciiArt("-*/---*-*//**", fontMap)
	fmt.Println(generatedArt)
}
