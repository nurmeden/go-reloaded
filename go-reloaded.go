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

	result_word := ""
	dat, _ := os.ReadFile("sample.txt")
	splittedString := strings.Split(string(dat), " ")
	fmt.Println(splittedString)
	for i := range splittedString {
		if string(splittedString[i]) == "(up)" {
			Upperstr := strings.ToUpper(splittedString[i-1])
			splittedString[i-1] = Upperstr
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(low)" {
			Lowerstr := strings.ToLower(splittedString[i-1])
			splittedString[i-1] = Lowerstr
			splittedString[i] = ""
		} else if string(splittedString[i]) == "(cap, 6)" {
			numstr := ""
			for _, value := range splittedString[i] {
				if '0' < value && value < '9' {
					numstr = string(value)
				}
			}
			fmt.Println(numstr)
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
		}
	}

	for i := 0; i < len(splittedString); i++ {
		if i != 0 {
			result_word += " " + string(splittedString[i])
		} else {
			result_word += string(splittedString[i])
		}
	}

	_, err2 := f.Write([]byte(result_word))
	check(err2)
}

func formatD(word string) bool {
}
