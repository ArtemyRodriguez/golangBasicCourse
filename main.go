package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println(">>> Калькулятор Индекса Массы тела <<<")
		userHeight, userWeight := getUserInput()
		IMT := calculateIMT(userHeight, userWeight)
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела!")
	case imt < 18.5:
		fmt.Println("У вас недостаток веса.")
	case imt < 25:
		fmt.Println("У вас нормальный вес!")
	case imt < 30:
		fmt.Println("У вас избыток массы тела!")
	default:
		fmt.Println("У вас ожирение!")
	}
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите повторить расчет(y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
