package main

import "fmt"

//Escreva uma função que receba um ponteiro para um booleano e altere o valor desse
//booleano para o seu inverso.

func inverso(x *bool) {
	*x = !*x
}
func main() {
	flag := true
	fmt.Println("antes da inversao:", flag)
	inverso(&flag)
	fmt.Println("depois da inversao:", flag)
}
