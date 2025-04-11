package modifier

import "strconv"

func processDigits(textArray []rune) []rune {
	sum, digitsCounter := 0, 0
	var digitInText int

	for index, value := range textArray {
		if value >= '0' && value <= '9' {
			digitsCounter++
			digitInText = int(value) - '0'
			sum += digitInText
			textArray = removeChars(textArray, index)
		}
	}

	if digitsCounter > 0 {
		textArray = append(textArray, ' ')
		stringSum := strconv.Itoa(sum)
		charsDigitsOfSum := textToCharArray(stringSum)

		for i, _ := range charsDigitsOfSum {
			textArray = append(textArray, charsDigitsOfSum[i])
		}
	}

	return textArray
}
