package main

import "fmt"

//Escreva uma função que receba um ponteiro para uma struct que contém informações
//de um produto (nome, preço e quantidade em estoque). A função deve atualizar o preço
//desse produto para um novo valor fornecido como argumento.

type Produto struct {
	Nome  string
	Preco float64
	qntd  int
}

func updatePreco(p *Produto, novoPreco float64) {
	p.Preco = novoPreco
}
func main() {
	produto := Produto{Nome: "short", Preco: 20.00, qntd: 5}
	fmt.Println("antes da atualização: ", produto)
	updatePreco(&produto, 60.0)
	fmt.Println("depois da atualização", produto)
}
