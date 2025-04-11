package modifier

func TextModifier(input string) string {

	textArray := textToCharArray(input)
	textWithoutSpaces := deleteSpaces(textArray)
	textChangedChars := changePlaces(textWithoutSpaces)
	textWithoutPluses := changePluses(textChangedChars)
	textWithProcessedDigits := processDigits(textWithoutPluses)

	return string(textWithProcessedDigits)
}

func removeChars(slice []rune, i int) []rune {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
