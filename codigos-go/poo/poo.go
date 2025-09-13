package main

import "fmt"

type Entregador struct {
	nome string
	entregas int
	meta int
}

func (f Entregador) meta_entrega() {
	if f.entregas >= f.meta {
		fmt.Println(f.nome, "Bateu a meta!")
	} else {
		fmt.Println(f.nome, "Nao bateu a meta!")
	}
}

func main() {
	entregadores := []Entregador{
		{"Ikaro", 20, 10}, 
		{"Jao", 10, 11},
		{"Clebiosmar", 500, 300},
	}
	Entregador.meta_entrega(entregadores[1])
}