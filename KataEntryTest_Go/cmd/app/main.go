package main

import (
	"fmt"
	"kata/entryTest/internal/modifier"
)

func main() {
	targetText := "Я люблю Golang всем сердцем! 25"
	input := "Я ю-лб8ю-л Gol7ang   вс6ем серд4цем+"
	fmt.Println("Исходный текст:       ", input)
	fmt.Println("Целевой текст:        ", targetText)

	modifiedText := modifier.TextModifier(input)
	printResult(modifiedText)

	checkResult(modifiedText, targetText)

}

func printResult(text string) {
	fmt.Println("Преобразованный текст: " + text)
}

func checkResult(modifiedText string, targetText string) {
	if modifiedText == targetText {
		fmt.Println("Преобразование выполнено успешно!!!")
	} else {
		fmt.Println("!!!!!       Нужно поработать еще!!!")
	}
}
