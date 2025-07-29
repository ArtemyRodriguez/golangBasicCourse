package main

import (
	"fmt"
	"math"
)

func main() {

	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println(">>> Калькулятор Индекса Массы тела <<<")
	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	outputResult(IMT)

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}
