package modifier

func textToCharArray(text string) []rune {
	var charArray []rune
	for _, char := range text {
		charArray = append(charArray, char)

	}
	return charArray

}
