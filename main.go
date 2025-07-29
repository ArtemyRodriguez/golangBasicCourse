package main

import (
	"fmt"
	"math"
)

func main() {

	const IMTPower = 2
	var userHeight float64
	var userWeight float64

	//fmt.Println(">>> Калькулятор Индекса Массы тела <<<\n")

	//многострочный метод
	fmt.Print(`>>> Калькулятор Индекса Массы тела <<<
	Введите ваш рост в сантиметрах: `)

	//fmt.Println(">>> Калькулятор Индекса Массы тела <<<")
	//fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)

	//fmt.Print(IMT)
	//fmt.Print("Ваш индекс массы тела: ", IMT)
	//fmt.Printf("Ваш индекс массы тела: %v", IMT)
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)

}
