package functions

func AddTxtExtension(filename string) string {
	var result []rune
	for _, char := range filename {
		result = append(result, char)
	}
	txtExtension := ".txt"
	for _, char := range txtExtension {
		result = append(result, char)
	}
	return string(result)
}
