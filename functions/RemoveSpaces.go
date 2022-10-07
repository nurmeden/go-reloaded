package functions

func RemoveSpaces(splittedString []string, count int) ([]string, int) {
	for i := 0; i < len(splittedString); i++ {
		if splittedString[i] == "" {
			splittedString = RemoveIndex(splittedString, i)
			count++
			i--
		}
	}
	return splittedString, count
}
