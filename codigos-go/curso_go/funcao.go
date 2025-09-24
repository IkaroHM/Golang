package main

import (
	"fmt"
)

func main() {
	res, str := somaa(10, 20)
	fmt.Println(res, str)
	resultado := somaaa(10, 20)
	fmt.Println(resultado)
	resultadoo := somatudo(10, 20, 30, 45)
	fmt.Println(resultadoo)
}

func somaa(a, b int) (int, string) {
	return a + b, "soma"
}

func somaaa(a, b int ) ( resultado int) {
	resultado = a + b
	return
}

func somatudo(x...int) int {
	resultadoo := 0
	for _, v := range x {
		resultadoo += v
	}
	return resultadoo
}