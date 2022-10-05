package functions

func RemoveSpaces(splittedString []string, count int) ([]string, int) {
	for i := 0; i < len(splittedString); i++ {
		if splittedString[i] == "''" {
			RemoveIndex(splittedString, i)
			count++
		}
	}
	return splittedString, count
}
