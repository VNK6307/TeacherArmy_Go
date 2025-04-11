package main

import (
	"bufio"
	"fmt"
	"kata/entryTest/internal/modifier"
	"os"
)

func main() {
	targetText := "Я люблю Golang всем сердцем! 25"
	input := "Я ю-лб8ю-л Gol7ang   вс6ем серд4цем+"
	fmt.Println("Исходный текст:       ", input)
	fmt.Println("Целевой текст:        ", targetText)

	modifiedText := modifier.TextModifier(input)
	printResult(modifiedText)

	checkResult(modifiedText, targetText)

	fmt.Println("Введите текст для преобразования: ")
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	input = scanner.Text()
	fmt.Println("Вы ввели фразу для преобразования:", input)

	modifiedText = modifier.TextModifier(input)
	printResult(modifiedText)

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

// генрих1  играет+2   л-июбит0школу
// генрих играет! илюбитшколу 3

// Введено   три пробела
// Введено   два раза     по     пять   пробелов. "
