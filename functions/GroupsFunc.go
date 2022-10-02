package functions

func GroupsFunc(splittedString []string) []string {
	for i := 0; i < len(splittedString); i++ {
		if (33 <= splittedString[i][0] && 47 >= splittedString[i][0]) || (58 <= splittedString[i][0] && 63 >= splittedString[i][0]) {
			splittedString[i-1] = splittedString[i-1] + splittedString[i]
			splittedString = append(splittedString[:i], splittedString[i+1:]...)

		} // for j := 0; j < len(splittedString[i]); j++ {
		// 	if (33 <= splittedString[i][j] && 47 >= splittedString[i][j]) || (58 <= splittedString[i][j] && 63 >= splittedString[i][j]) {
		// 		splittedString[i-1] += string(splittedString[i][j])
		// 		splittedString[i] = RemoveIndexstr(splittedString[i], j)
		// 	} else {
		// 		break
		// 	}
		// 	// splittedString[i] = strings.Replace(splittedString[i], string(splittedString[i][j]), "", -1)
		// }
	}
	return splittedString
}
