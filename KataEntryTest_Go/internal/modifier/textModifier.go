package modifier

func TextModifier(input string) string {

	modifiedText := processDigits(
		changePluses(
			changePlaces(
				deleteSpaces(
					textToCharArray(input)))))

	return string(modifiedText)

}

func removeChars(slice []rune, i int) []rune {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
