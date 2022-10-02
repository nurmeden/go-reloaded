package functions

import "strconv"

func replaceDigits(splittedString []string, digit int, i int, count int) ([]string, int) {
	splittedString[i-1] = strconv.Itoa(digit)
	RemoveIndex(splittedString, i)
	count++
	return splittedString, count
}
