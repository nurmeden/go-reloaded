package functions

import "strings"

func CheckVowels(splittedString []string, count int) []string {
	Vowels := "aeiouhAEIOUH"
	for i := 0; i < len(splittedString); i++ {
		if splittedString[i] == "a" {
			if strings.HasPrefix(Vowels, splittedString[i]) {
				splittedString[i] = "an"
			}
		}
	}
	return splittedString
}
