package functions

func CheckPunctuationMark(splittedString []string, count int, flag bool) []string {
	countmarks := 1
	for i := 0; i < len(splittedString); i++ {
		if splittedString[i] == "'" {
			if countmarks%2 != 0 {
				splittedString[i+1] = splittedString[i] + splittedString[i+1]
				RemoveIndex(splittedString, i)
				countmarks++
			} else {
				splittedString[i-1] += splittedString[i]
				RemoveIndex(splittedString, i)
				countmarks++
			}
		}
	}
	return splittedString
}
