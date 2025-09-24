package main

import "fmt"
import "errors"

func soma(a float64, b float64) (float64, error) {
	return a+b, errors.New("nada")
}

func main() {
	if retorno, err := soma(10, 20); err != nil {
		fmt.Println(retorno, err)
	}
}