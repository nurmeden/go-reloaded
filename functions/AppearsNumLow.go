package functions

import "strings"

func AppearsNumLow(splittedString []string, i int, numm int, count int, flag bool) ([]string, int, bool) {
	for j := 0; j < len(splittedString[i-numm:i]); j++ {
		Lowerstr := strings.ToLower(splittedString[(i+j)-numm])
		splittedString[(i+j)-numm] = Lowerstr
	}
	splittedString = RemoveIndex(splittedString, i)
	splittedString = RemoveIndex(splittedString, i)
	count += 2

	// if splittedString[i+1][len(splittedString[i+1])-1] == ',' {

	// }

	if flag == true {
		splittedString, count, flag = Check(splittedString, count, flag)
	}
	return splittedString, count, flag
}
