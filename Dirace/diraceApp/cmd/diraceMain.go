package main

import (
	"dirace/internal/raceControl"
	"fmt"
)

func main() {
	fmt.Println("Добро пожаловать, $racer!\n    Выберите предпочтительный способ авторизации:")

	raceControl.ShowRaceClass("Standard-90")

}
