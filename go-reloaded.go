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
			i, _ := strconv.Atoi(splittedString[i-1])
			Binstr := strconv.FormatInt(int64(i), 10)
			splittedString[i-1] = Binstr
			fmt.Println(splittedString[i-1])
			splittedString[i] = ""
		}
	}

	for i := 0; i < len(splittedString); i++ {
		_, err2 := f.Write([]byte(splittedString[i]))
		check(err2)
	}
}
