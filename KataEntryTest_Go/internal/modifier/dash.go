package modifier

import "fmt"

func changePlaces(textArray []rune) []rune {
	if textArray[0] == '-' || textArray[len(textArray)-1] == '-' {
		fmt.Println("Невозможно поменять символы местами, так как символ '-' является первым или последним элементом строки.")
	} else {
		for i := 0; i < (len(textArray) - 1); i++ {
			if textArray[i] == '-' {
				charBefore := textArray[i-1]
				charAfter := textArray[i+1]
				textArray[i-1] = charAfter
				textArray[i+1] = charBefore
				textArray = removeChars(textArray, i)
			}
		}
	}

	return textArray
}
