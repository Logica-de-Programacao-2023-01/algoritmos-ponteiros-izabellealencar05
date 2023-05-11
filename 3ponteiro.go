package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro para uma string
// e atualize o valor da string para seu reverso.

func inverso(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}

func main() {
	str := "izabelle"
	fmt.Print("antes da inversão: ", str)
	inverso(&str)
	fmt.Print("após a inversão: ", str)
}
