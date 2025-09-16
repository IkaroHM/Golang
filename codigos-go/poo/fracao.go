package main

import (
	"fmt"
)

var (
	Numr int
	Denr int
)

func Somafracao(a, b, c, d int){
	if b != d {
		Numr = a * d + b * c
		Denr = b * d
	} else if b == d {
		Numr = a + c
		Denr = b
	}
	fmt.Printf("%d/%d somado por %d/%d é igual a %.d/%d\n", a, b, c, d, Numr, Denr)
}

func Subtraifracao(a, b, c, d int){
	if b != d {
		Numr = a * d - b * c
		Denr = b * d
	} else {
		Numr = a - c
		Denr = b
	}
	fmt.Printf("%d/%d subtraido por %d/%d é igual a %.d/%d\n", a, b, c, d, Numr, Denr)
}

func Multiplicafracao(a, b, c, d int){
	Numr = a * c
	Denr = b * d
	fmt.Printf("%d/%d multiplicado por %d/%d é igual a %.d/%d\n", a, b, c, d, Numr, Denr)
}

func dividirfracao(a, b, c, d int){
	Numr = a * d
	Denr = b * c
	fmt.Printf("%d/%d dividido por %d/%d é igual a %.d/%d\n", a, b, c, d, Numr, Denr)
}

func main() {
	Multiplicafracao(2, 3, 2, 7)
	dividirfracao(3, 4, 5, 2)
}