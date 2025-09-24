package main

import (
	"fmt"

	"github.com/IkaroHM/codigos-go/curso_go/math"
	"github.com/google/uuid"
)

func main() {
	a := 10
	b := "World"
	c := 3.14
	d := false
	fmt.Printf("%v %v  %v  %v ", a, b, c, d)
	resultado := matematica.Soma (1, 1)
	resultado2 := matematica.A
	fmt.Printf("%v %v", resultado, resultado2)
	novoUuid := uuid.New()
	novoUuidstr := novoUuid.String()
	fmt.Println("o seu UUid é :", novoUuidstr)
}