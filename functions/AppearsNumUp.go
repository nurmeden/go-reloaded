package functions

import "strings"

func AppearsNumUp(splittedString []string, i int, numm int, count int, flag bool) ([]string, int, bool) {
	for j := 0; j < len(splittedString[i-numm:i]); j++ {
		Lowerstr := strings.ToUpper(splittedString[(i+j)-numm])
		splittedString[(i+j)-numm] = Lowerstr
	}
	splittedString = RemoveIndex(splittedString, i)
	splittedString = RemoveIndex(splittedString, i)
	count += 2
	if flag == true {
		splittedString, count, flag = Check(splittedString, count, flag)
	}
	return splittedString, count, flag
}
