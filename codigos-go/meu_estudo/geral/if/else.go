package main

import "fmt"
import "errors"

func oi() (string, error) {
	return "teste", errors.New("test")
}

func drible(n string) string {
	if n == "1" {
		return "chapeu"
	} else if n == "2" {
		return "caneta"
	} else {
		return "pedalada"
	}
}

func main() {

	if retorno, err := oi(); err != nil {
		fmt.Println("true")
		fmt.Println(retorno, err)
	}

	if returno := drible("1"); returno == "pedalada" {
		fmt.Println("true")
	}

}
