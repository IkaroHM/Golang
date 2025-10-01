package main

import (
	"fmt"
)

type Carru struct{
	Nome string
}

func (c Carru) andou() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carru{
		Nome: "Uno Com Escada",
	}

	carro.andou()

	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)
}

func abc(a *int) {
	*a = 200
}