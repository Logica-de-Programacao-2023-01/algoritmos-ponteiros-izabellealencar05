package main

import "fmt"

//Escreva uma função swap que receba dois ponteiros para int como argumentos e troque
//os valores apontados por eles.Escreva uma função swap que
//receba dois ponteiros para int como argumentos e troque
//os valores apontados por eles.

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	a := 5
	b := 10
	fmt.Println("antes da troca: a =", a, "e b =", b)
	swap(&a, &b)
	fmt.Println("depois da troca: a =", a, "e b =", b)
}
