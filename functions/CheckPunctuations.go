package functions

import (
	"strings"
)

func CheckPunctuations(splittedString []string, count int) ([]string, int) {
	for i := 0; i < len(splittedString[:len(splittedString)-count]); i++ {
		if splittedString[i] == "," || splittedString[i] == "." || splittedString[i] == "?" || splittedString[i] == "!" || splittedString[i] == ";" || splittedString[i] == ":" { // IT,  WAS
			splittedString[i-1] += splittedString[i] // []string{"was,",it"}
			RemoveIndex(splittedString, i)
			count++
		} else if splittedString[i][0] == ',' || splittedString[i][0] == '.' || splittedString[i][0] == '?' || splittedString[i][0] == '!' || splittedString[i][0] == ';' || splittedString[i][0] == ':' { // IT, WAS
			splittedString[i-1] += string(splittedString[i][0]) // hello(0) ,world(1) of,the
			RemoveIndex(splittedString, i)
			count++
		}
		k := strings.IndexRune(splittedString[i], ',')
		las := strings.HasSuffix(splittedString[i], ",")
		if k >= 0 && len(splittedString[i]) > 1 && !las {
			splitstr := strings.Split(splittedString[i], ",")
			splitstr[0] += "," + " " + splitstr[1]
			splittedString[i] = splitstr[0]
			count += 2
		}
	}
	return splittedString, count
}
