package functions

import "strings"

func AppearsNumCap(splittedString []string, i int, numm int, count int, flag bool) ([]string, int, bool) {
	for j := 0; j < len(splittedString[i-numm:i]); j++ {
		Titlestr := strings.Title(splittedString[(i+j)-numm])
		splittedString[(i+j)-numm] = Titlestr
	}
	RemoveIndex(splittedString, i)
	RemoveIndex(splittedString, i)
	count += 2
	if flag == true {
		splittedString, count, flag = Check(splittedString, count, flag)
	}
	return splittedString, count, flag
}
