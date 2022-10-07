package functions

func FirstArgs(splittedString []string) []string {
	count := 0
	args := []string{"(cap)", "(up)", "(low)"}
	args2 := []string{"(cap,", "(up,", "(low,"}
	for i := 0; i < len(args); i++ {
		if splittedString[0] == args[i] {
			splittedString[0] = ""
			splittedString, count = RemoveSpaces(splittedString, count)
			if len(splittedString) != 0 {
				splittedString = FirstArgs(splittedString)
				break
			} else {
				break
			}
		}
	}
	for i := 0; i < len(args2); i++ {
		if splittedString[0] == args2[i] {
			splittedString[0] = ""
			splittedString[1] = ""
			splittedString, count = RemoveSpaces(splittedString, count)
			if len(splittedString) != 0 {
				splittedString = FirstArgs(splittedString)
				break
			} else {
				break
			}
		}
	}
	return splittedString
}
