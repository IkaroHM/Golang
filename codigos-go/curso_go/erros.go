package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	res, err := soma(4, 2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res, err)
}

func soma(x, y int) ( int, error){
	res := x + y

	if res>10 {
		return 0, errors.New("O total é maior que 10")
	}
	return res, nil
}

