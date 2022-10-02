package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func Check(splittedString []string, count int, flag bool) ([]string, int, bool) {
	// for i := 0; i < len(splittedString); i++ {
	// 	fmt.Print("[" + splittedString[i] + "]")
	// }
	for i := range splittedString {
		if string(splittedString[i]) == "(cap)" {
			Titlestr := strings.Title(splittedString[i-1])
			splittedString, count, flag = replace(splittedString, Titlestr, i, count, flag)
		} else if string(splittedString[i]) == "(up)" {
			Upperstr := strings.ToUpper(splittedString[i-1])
			splittedString, count, flag = replace(splittedString, Upperstr, i, count, flag)
		} else if string(splittedString[i]) == "(low)" {
			Lowerstr := strings.ToLower(splittedString[i-1])
			splittedString, count, flag = replace(splittedString, Lowerstr, i, count, flag)
		} else if string(splittedString[i]) == "(bin)" {
			Decstr, err := strconv.ParseInt(string(splittedString[i-1]), 2, 64)
			CheckErr(err)
			splittedString, count = replaceDigits(splittedString, int(Decstr), i, count)
		} else if string(splittedString[i]) == "(hex)" {
			Hexstr, err := strconv.ParseInt(string(splittedString[i-1]), 16, 64)
			CheckErr(err)
			splittedString, count = replaceDigits(splittedString, int(Hexstr), i, count)
		} else if string(splittedString[i]) == "(cap," {
			if splittedString[i+1][len(splittedString[i+1])-1] == ')' {
				num := splittedString[i+1]
				numm, err := strconv.Atoi(string(num[:len(num)-1]))
				CheckErr(err)
				if numm <= 0 {
					flag = false
					fmt.Println("a number is negative or zero")
				} else if numm <= len(splittedString[:i]) {
					splittedString, count, flag = AppearsNumCap(splittedString, i, numm, count, flag)
				} else if numm >= len(splittedString[:i]) {
					numm = len(splittedString[:i])
					splittedString, count, flag = AppearsNumCap(splittedString, i, numm, count, flag)
				}
			}
		} else if string(splittedString[i]) == "(low," {
			if splittedString[i+1][len(splittedString[i+1])-1] == ')' {
				num := splittedString[i+1] // 43)
				numm, err := strconv.Atoi(string(num[:len(num)-1]))
				CheckErr(err)
				if numm <= 0 {
					flag = false
					fmt.Println("a number is negative or zero")
					RemoveIndex(splittedString, i)
					RemoveIndex(splittedString, i)
					count += 2
				} else if numm <= len(splittedString[:i]) && numm > 0 {
					splittedString, count, flag = AppearsNumLow(splittedString, i, numm, count, flag)
				} else if numm >= len(splittedString[:i]) && numm > 0 {
					numm = len(splittedString[:i])
					splittedString, count, flag = AppearsNumLow(splittedString, i, numm, count, flag)
				}
			}
		} else if string(splittedString[i]) == "(up," {
			if splittedString[i+1][len(splittedString[i+1])-1] == ')' {
				num := splittedString[i+1]
				numm, err := strconv.Atoi(string(num[:len(num)-1]))
				CheckErr(err)
				if numm <= 0 {
					flag = false
					fmt.Println("a number is negative or zero")
				} else if numm <= len(splittedString[:i]) && numm > 0 {
					splittedString, count, flag = AppearsNumUp(splittedString, i, numm, count, flag)
				} else if numm >= len(splittedString[:i]) && numm > 0 {
					numm = len(splittedString[:i])
					splittedString, count, flag = AppearsNumUp(splittedString, i, numm, count, flag)
				}
			}
		} else if string(splittedString[i]) == "" {
			RemoveIndex(splittedString, i)
			count++
		}
	}

	return splittedString, count, flag
}
