package functions

func replace(splittedString []string, str string, i int, count int, flag bool) ([]string, int, bool) {
	if i != len(splittedString)-1 {
		splittedString[i-1] = str
		splittedString = RemoveIndex(splittedString, i)
		count++
		splittedString, count, flag = Check(splittedString, count, flag)
	} else {
		splittedString[i-1] = str
		splittedString = RemoveIndex(splittedString, i)
		count++
	}
	return splittedString, count, flag
}
