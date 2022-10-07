package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {
	flag := true

	if len(os.Args) == 3 && (os.Args[1] == "sample.txt" && os.Args[2][len(os.Args[2])-4:] == ".txt") {
		count := 0
		dat, err := os.ReadFile("sample.txt")
		functions.CheckErr(err)
		str := string(dat)
		splittedString := strings.Split(str, " ")

		splittedString = functions.FirstArgs(splittedString)

		splittedString, count = functions.RemoveSpaces(splittedString, count)

		splittedString, count = functions.CheckPunctuationMark(splittedString, count, flag)
		// splittedString = functions.GroupsFunc(splittedString)

		splittedString, count, flag = functions.Check(splittedString, count, flag)

		splittedString = functions.GroupsFunc(splittedString)

		splittedString = functions.CheckPunctuations(splittedString)

		splittedString = functions.CheckVowels(splittedString, count)

		result := functions.Result_words(splittedString, count)
		resultbyte := []byte(result)

		errResult := os.WriteFile(os.Args[2], resultbyte, 0644)
		functions.CheckErr(errResult)
	} else {
		fmt.Println("error arguments")
	}

	// str := [3]string{"up","low","cap"}
	// f, err := os.Create(os.Args[2])
	// functions.CheckErr(err)

	// defer f.Close()

	// _, err2 := f.Write([]byte(result))
	// functions.CheckErr(err2)
}
