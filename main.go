package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	//циклы
	for i := 0; i < 10; i++ {
		if i == 5 {
			//break
			continue
		}
		fmt.Printf("%d\n", i)
	}
	fmt.Println(">>> Калькулятор Индекса Массы тела <<<")
	userHeight, userWeight := getUserInput()
	IMT := calculateIMT(userHeight, userWeight)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела!")
	case IMT < 18.5:
		fmt.Println("У вас недостаток веса.")
	case IMT < 25:
		fmt.Println("У вас нормальный вес!")
	case IMT < 30:
		fmt.Println("У вас избыток массы тела!")
	default:
		fmt.Println("У вас ожирение!")
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
