package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("result.txt")
	check(err)

	defer f.Close()
	result_words := ""
	dat, _ := os.ReadFile("sample.txt")
	splittedString := strings.Split(string(dat), " ")
	for i := range splittedString {
		if string(splittedString[i]) == "(up)" {
			Upperstr := strings.ToUpper(splittedString[i-1])
			splittedString[i-1] = Upperstr
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(low)" {
			Lowerstr := strings.ToLower(splittedString[i-1])
			splittedString[i-1] = Lowerstr
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(cap)" {
			Titlestr := strings.Title(splittedString[i-1])
			splittedString[i-1] = Titlestr
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(bin)" {
			Decstr, _ := strconv.ParseInt(string(splittedString[i-1]), 2, 64)
			splittedString[i-1] = strconv.Itoa(int(Decstr))
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(hex)" {
			Hexstr, _ := strconv.ParseInt(string(splittedString[i-1]), 16, 64)
			splittedString[i-1] = strconv.Itoa(int(Hexstr))
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(cap," {
			num := splittedString[i+1]
			numm, _ := strconv.Atoi(string(num[:len(num)-1]))
			for j := 0; j < len(splittedString[i-numm:i]); j++ {
				Titlestr := strings.Title(splittedString[(i+j)-numm])
				splittedString[(i+j)-numm] = Titlestr
			}
			splittedString[i] = ""
			splittedString[i+1] = ""
		} else if string(splittedString[i]) == "(low," {
			num := splittedString[i+1]
			numm, _ := strconv.Atoi(string(num[0]))
			for j := 0; j < len(splittedString[i-numm:i]); j++ {
				Lowerstr := strings.ToLower(splittedString[(i+j)-numm])
				splittedString[(i+j)-numm] = Lowerstr
			}
			splittedString[i] = ""
			splittedString[i+1] = ""
		} else if string(splittedString[i]) == "(up," {
			num := splittedString[i+1]
			numm, _ := strconv.Atoi(string(num[0]))
			for j := 0; j < len(splittedString[i-numm:i]); j++ {
				Lowerstr := strings.ToUpper(splittedString[(i+j)-numm])
				splittedString[(i+j)-numm] = Lowerstr
			}
			splittedString[i] = ""
			splittedString[i+1] = ""
		}
	}
	fmt.Println(splittedString)

	for i := 0; i < len(splittedString); i++ {
		if i != 0 {
			result_words += " " + string(splittedString[i])
		} else {
			result_words += string(splittedString[i])
		}
	}

	fmt.Println(result_words)
	_, err2 := f.Write([]byte(result_words))
	check(err2)
}

// func splittedString(dat string) []string {
// 	splittedString := strings.Split(dat, " ")
// 	fmt.Println(splittedString)
// 	for i := range splittedString {
// 		if string(splittedString[i]) == "(up)" {
// 			Upperstr := strings.ToUpper(splittedString[i-1])
// 			splittedString[i-1] = Upperstr
// 			splittedString[i] = ""
// 		} else if string(splittedString[i]) == "(low)" {
// 			Lowerstr := strings.ToLower(splittedString[i-1])
// 			splittedString[i-1] = Lowerstr
// 			splittedString[i] = ""
// 		} else if string(splittedString[i]) == "(cap, 6)" {
// 			Titlestr := strings.Title(splittedString[i-1])
// 			splittedString[i-1] = Titlestr
// 			splittedString[i] = ""
// 		} else if string(splittedString[i]) == "(bin)" {
// 			Decstr, _ := strconv.ParseInt(string(splittedString[i-1]), 2, 64)
// 			splittedString[i-1] = strconv.Itoa(int(Decstr))
// 			splittedString[i] = ""
// 		} else if string(splittedString[i]) == "(hex)" {
// 			Hexstr, _ := strconv.ParseInt(string(splittedString[i-1]), 16, 64)
// 			splittedString[i-1] = strconv.Itoa(int(Hexstr))
// 			splittedString[i] = ""
// 		} else if string(splittedString[i]) == "(up," {
// 			num := splittedString[i+1]
// 			fmt.Println(num)
// 		}
// 	}
// 	return splittedString
// }

// func result_word(splittedString []string) string {
// 	result_words := ""
// 	for i := 0; i < len(splittedString); i++ {
// 		if i != 0 {
// 			result_words += " " + string(splittedString[i])
// 		} else {
// 			result_words += string(splittedString[i])
// 		}
// 	}
// 	return result_words
// }
