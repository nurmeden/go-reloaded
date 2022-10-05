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

	os.Args = []string{".", "sample.txt", "result.txt"}

	if len(os.Args) == 3 && os.Args[1] == "sample.txt" && os.Args[2] == "result.txt" {
		count := 0
		dat, _ := os.ReadFile("sample.txt")

		splittedString := strings.Split(string(dat), " ")

		splittedString, count, flag = functions.Check(splittedString, count, flag)
		fmt.Println(splittedString)

		splittedString = functions.GroupsFunc(splittedString)
		fmt.Println(splittedString)

		splittedString, count = functions.CheckPunctuationMark(splittedString, count, flag)
		fmt.Println(splittedString)

		splittedString = functions.CheckPunctuations(splittedString)
		fmt.Println(splittedString)

		splittedString = functions.CheckVowels(splittedString, count)
		fmt.Println(splittedString)

		splittedString, count = functions.RemoveSpaces(splittedString, count)
		fmt.Println(splittedString)

		result := functions.Result_words(splittedString, count)

		_, err2 := f.Write([]byte(result))
		functions.CheckErr(err2)
	} else {
		fmt.Println("error arguments")
	}


	// str := [3]string{"up","low","cap"}
}
