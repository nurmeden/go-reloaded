package functions

func GroupsFunc(splittedString []string) []string {
	for i := 0; i < len(splittedString); i++ {
		// if i == len(splittedString)-1 && (33 <= splittedString[i][0] && 38 >= splittedString[i][0]) || (42 <= splittedString[i][0] && 47 >= splittedString[i][0]) || (58 <= splittedString[i][0] && 63 >= splittedString[i][0]) {
		// 	splittedString[i-1] = splittedString[i-1] + splittedString[i]
		// 	splittedString = append(splittedString[:i], splittedString[i+1:]...)
		// }
		if (33 <= splittedString[i][0] && 38 >= splittedString[i][0]) || (42 <= splittedString[i][0] && 47 >= splittedString[i][0]) || (58 <= splittedString[i][0] && 63 >= splittedString[i][0]) {
			splittedString[i-1] = splittedString[i-1] + splittedString[i]
			splittedString = append(splittedString[:i], splittedString[i+1:]...)
		}
	}
	return splittedString
}
