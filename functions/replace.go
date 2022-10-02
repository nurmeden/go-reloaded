package functions

func replace(splittedString []string, str string, i int, count int, flag bool) ([]string, int, bool) {
	splittedString[i-1] = str
	RemoveIndex(splittedString, i)
	count++
	if flag == true {
		splittedString, count, flag = Check(splittedString, count, flag)
	}
	return splittedString, count, flag
}
