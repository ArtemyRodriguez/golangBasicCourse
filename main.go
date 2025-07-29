package main

import (
	"fmt"
	"math"
)

func main() {

	var userHeight float64
	var userWeight float64
	fmt.Println(">>> Калькулятор Индекса Массы тела <<<")
	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userWeight)
	IMT := calculateIMT(userHeight, userWeight)
	outputResult(IMT)

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userHeight float64, userWeight float64) float64 {
	const IMTPower = 2
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT
}
