package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("result.txt")
	functions.CheckErr(err)

	defer f.Close()

	flag := true

	if len(os.Args) == 3 && os.Args[1] == "sample.txt" && os.Args[2] == "result.txt" {
		count := 0
		dat, _ := os.ReadFile("sample.txt")

		splittedString := strings.Split(string(dat), " ")

		splittedString, count, flag = functions.Check(splittedString, count, flag)

		splittedString = functions.GroupsFunc(splittedString)

		splittedString = functions.CheckPunctuationMark(splittedString, count, flag)

		splittedString = functions.CheckPunctuations(splittedString)

		splittedString = functions.CheckVowels(splittedString, count)

		result := functions.Result_words(splittedString, count)

		_, err2 := f.Write([]byte(result))
		functions.CheckErr(err2)
	} else {
		fmt.Println("error arguments")
	}

	// str := [3]string{"up","low","cap"}
}
