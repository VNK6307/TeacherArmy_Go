package modifier

func changePluses(textArray []rune) []rune {
	for index, value := range textArray {
		if value == '+' {
			textArray[index] = '!'
		}
	}
	return textArray
}
