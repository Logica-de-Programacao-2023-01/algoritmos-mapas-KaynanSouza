package main

import "fmt"

// Escreva uma função que receba um mapa com valores
// inteiros e retorne a soma de todos os valores.
func main() {
	numero := map[string]int{
		"numero1": 20,
		"numero2": 30,
		"numero3": 10,
	}
	resp := soma(numero)
	fmt.Print(resp)
}
func soma(numero map[string]int) int {
	var soma int
	for _, num := range numero {
		soma += num
	}
	return soma
}
