package main

import (
	"fmt"
)

// Escreva uma função que receba dois mapas e retorne um novo
// mapa contendo todos os elementos dos mapas de entrada.
// Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.
func main() {
	um := map[string]int{
		"a": 2,
		"b": 3,
	}
	dois := map[string]int{
		"c": 4,
		"d": 5,
	}
	resp := soma(dois, um)
	for chave, valor := range resp {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}
func soma(um, dois map[string]int) map[string]int {
	tres := make(map[string]int)
	for chave, valor := range um {
		tres[chave] = valor
	}
	for chave, valor := range dois {
		tres[chave] = valor
	}

	return tres
}
