package main

import "fmt"

// Escreva uma função que conte a ocorrência de cada palavra em uma string
// e retorne um mapa onde as chaves são as palavras encontradas e os valores
// são o número de ocorrências de cada palavra.
func main() {
	str := "ola,mundo"
	resp := oco(str)
	for char, cont := range resp {
		fmt.Printf("%c: %d\n", char, cont)
	}
}
func oco(str string) map[rune]int {
	occurrences := make(map[rune]int)
	for _, char := range str {
		occurrences[char]++
	}
	return occurrences
}
