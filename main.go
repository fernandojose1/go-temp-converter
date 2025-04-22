package main

import (
	"fmt"
)

func main() {
	var temp float64
	var unit string

	fmt.Println("--- Conversor de temperatura ---")
	fmt.Println("Este programa converte temperaturas entre Celsius e Fahrenheit.")

	fmt.Println("Digite a temperatura:")
	fmt.Scanln(&temp)

	fmt.Println("Digite a unidade (C para Celsius ou F para Fahrenheit):")
	fmt.Scanln(&unit)

	if unit == "C" || unit == "c" {
		fahrenheit := (temp * 9 / 5) + 32
		fmt.Printf("%.2f Celsius é igual a %.2f Fahrenheit\n", temp, fahrenheit)
	} else if unit == "F" || unit == "f" {
		celsius := (temp - 32) * 5 / 9
		fmt.Printf("%.2f Fahrenheit é igual a %.2f Celsius\n", temp, celsius)
	} else {
		fmt.Println("Unidade inválida. Por favor, insira C para Celsius ou F para Fahrenheit.")
	}
}
