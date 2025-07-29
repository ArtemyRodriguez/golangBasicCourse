package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Print(">>> Калькулятор Индекса Массы тела <<<\n")
	fmt.Print("Введите ваш рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
