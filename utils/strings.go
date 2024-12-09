package utils

func RepeatChar(char string, length int) []string {
	result := []string{}
	for range length {
		result = append(result, char)
	}
	return result
}
