package modifier

func deleteSpaces(textArray []rune) []rune {
	var arrayWithoutSpaces []rune
	for i := 0; i < len(textArray); i++ {
		arrayWithoutSpaces = append(arrayWithoutSpaces, textArray[i])
		if textArray[i] == ' ' {
			for j := i + 1; j < len(textArray); j++ {
				if textArray[j] == ' ' {
					continue
				}
				arrayWithoutSpaces = append(arrayWithoutSpaces, textArray[j])
				i = j
				break
			}
		}

	}
	return arrayWithoutSpaces
}
