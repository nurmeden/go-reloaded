package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	result := ""
	dat, _ := os.ReadFile("sample.txt")
	splittedString := strings.Split(string(dat), " ")
	for i, _ := range splittedString {
		if string(splittedString[i]) == "(up)" {
			Upperstr := strings.ToUpper(splittedString[i-1])
			if i != 0 {
				result = result + " " + Upperstr
			}
			i++
		} else if string(splittedString[i]) == "(low, 3)" {
			Lowerstr := strings.ToLower(splittedString[i-1])
			if i != 0 {
				result = result + " " + Lowerstr
			}
			i++
		}
		if i != 0 {
			result = result + " " + splittedString[i]
		} else {
			result = result + splittedString[i]
		}
	}

	_, err2 := f.Write([]byte(result))
	check(err2)
	fmt.Println(" ")
	fmt.Println(splittedString[2])
}
