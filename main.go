package main

import "fmt"

func main() {
	const USDoEUR = 0.85
	const USDoRUB = 82.13
	const EURoRUB = USDoRUB / USDoEUR

    fmt.Printf("1 USD = %.2f EUR\n", USDoEUR)
    fmt.Printf("1 USD = %.2f RUB\n", USDoRUB)
    fmt.Printf("1 EUR = %.2f RUB\n", EURoRUB)
}