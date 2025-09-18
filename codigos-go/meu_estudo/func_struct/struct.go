package main

import "fmt"

func main() {
	var user user = user {
		name: "Ikaro | ",
		age: 20,
		teste2: " | teste2",
		Email: " | email | ",
	}

	fmt.Println(user)
}

/*
letra minuscula é privado
letra maiuscula é publico
*/

type user struct {
	name string
	age int
	teste2 string
	Email string
	senha string
}