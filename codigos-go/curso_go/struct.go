package main

import (
	"fmt"
)

type Cliente struct {
	Nome string
	Email string
	Cpf string
}

func (c Cliente) Exibir() {
	fmt.Println("Exibindo cliente:", c.Nome)
}

type Clienteint struct {
	Cliente
	Pais string
}

func main() {
	cliente := Cliente{
		"Ikaro",
		"Ikaro@gmail.com",
		"97867633424",
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Bella", "Bella@gmail.com", "1432432424"}
	fmt.Printf("Nome: %v | email: %v | cpf: %v\n", cliente2.Nome, cliente2.Email, cliente2.Cpf)

	cliente3 := Clienteint{
		Cliente{
		"jao",
		"jao@gmail.com",
		"23423432",},
		"EUA",
	}
	fmt.Printf("Nome: %v | email: %v | cpf: %v\n", cliente3.Nome, cliente3.Email, cliente3.Cpf, cliente3.Pais)

	cliente.Exibir()
	cliente2.Exibir()
	cliente3.Exibir()
}