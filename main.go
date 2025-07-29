package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println(">>> Калькулятор Индекса Массы тела <<<")
	userHeight, userWeight := getUserInput()
	IMT := calculateIMT(userHeight, userWeight)
	if IMT < 16 {
		fmt.Println("У вас сильный дефицит массы тела!")
	} else if IMT < 18.5 {
		fmt.Print("У вас недостаток веса.")
	} else if IMT < 25 {
		fmt.Print("У вас нормальный вес.")
	} else if IMT < 30 {
		fmt.Print("У вас избыточная масса тела!")
	} else {
		fmt.Print("У вас ожирение!")
	}
	outputResult(IMT)

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userHeight float64, userWeight float64) float64 {
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
