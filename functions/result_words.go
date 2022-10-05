package functions

func Result_words(splittedString []string, count int) string {
	result_words := ""
	for i := 0; i < len(splittedString); i++ {
		if i != 0 {
			result_words += " " + string(splittedString[i])
		} else {
			result_words += string(splittedString[i])
		}
	}
	return result_words
}
