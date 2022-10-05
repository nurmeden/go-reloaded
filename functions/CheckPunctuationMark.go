package functions

func CheckPunctuationMark(splittedString []string, count int, flag bool) ([]string, int) {
	countmarks := 1
	for i := 0; i < len(splittedString)-1; i++ {
		for j := 0; j < len(splittedString[i]); j++ {
			if string(splittedString[i][j]) == "'" {
				if countmarks%2 != 0 {
					splittedString[i+1] = string(splittedString[i][j]) + splittedString[i+1]
					// splittedString[i] = splittedString[i][j+1:]
					if j == len(splittedString[i])-1 {
						RemoveIndex(splittedString, i)
					}
					countmarks++
				} else {
					splittedString[i-1] += string(splittedString[i][j])
					// splittedString[i] = splittedString[i][j+1:]
					if j == len(splittedString[i])-1 {
						RemoveIndex(splittedString, i)
					}
					countmarks++
				}
			} else {
				break
			}
		}
	}
	return splittedString, count
}
