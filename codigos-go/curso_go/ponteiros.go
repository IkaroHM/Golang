package main

import (
	"fmt"
)

func main() {

	//mem-endereco(10) <------ A <----- 10
	a := 10

	fmt.Println(&a)

	var ponteiro *int = &a
	fmt.Println(*ponteiro)
	*ponteiro = 50
	fmt.Println(*ponteiro)
}