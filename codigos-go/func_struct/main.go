package main

import "fmt"

var (
	nome string
	n1 int
	n2 int32
)

func main() {
	mensagem := "a"
	fmt.Println(mensagem)

	var b, f, s = true, 3.14, "sou uma string"
	fmt.Println(b, f, s)

	var total int
	total ++
	fmt.Println("Total:", total)

	nome = "Ikaro Henrique Martins"
	fmt.Println("Ola,", nome,"!")

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}