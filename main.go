package main

import "fmt"

func main() {
	const USDoEUR = 0.85
	const USDoRUB = 82.13
	const EURoRUB = USDoRUB / USDoEUR

	fmt.Printf("1 USD = %.2f EUR\n", USDoEUR)
	fmt.Printf("1 USD = %.2f RUB\n", USDoRUB)
	fmt.Printf("1 EUR = %.2f RUB\n", EURoRUB)

	// Вызов функции для считывания ввода
	userInput := getUserInput()
	fmt.Printf("Вы ввели: %.2f\n", userInput)
}

func getUserInput() float64 {
	var input float64
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&input)
	return input
}

// Пустая функция расчета конвертации валют
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	// TODO: здесь будет логика конвертации
	return 0
}
