package functions

func CheckPunctuations(splittedString []string) []string {
	for i := 0; i < len(splittedString)-1; i++ {
		arr := []string{}
		for j := 0; j < len(splittedString[i])-1; j++ {
			if ((splittedString[i][j+1] <= 90 && splittedString[i][j+1] >= 65) || (splittedString[i][j+1] >= 97 && splittedString[i][j+1] <= 122)) && (33 <= splittedString[i][j] && 47 >= splittedString[i][j]) || (58 <= splittedString[i][j] && 63 >= splittedString[i][j]) {
				first_half := splittedString[i][:j+1]
				second_falf := splittedString[i][j+1:]
				last_half := splittedString[i+1:]
				arr = append(arr, splittedString[:i]...)

				arr = append(arr, first_half)

				arr = append(arr, second_falf)

				arr = append(arr, last_half...)
				splittedString = arr
			}
		}
	}
	return splittedString
}
