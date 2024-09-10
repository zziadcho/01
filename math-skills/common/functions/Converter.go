package functions

func Converter(arr []byte) []string {
	var result []string

	for i, _ := range arr {
		result = append(result, string(arr[i] - '0'))
	}
	return result
}
