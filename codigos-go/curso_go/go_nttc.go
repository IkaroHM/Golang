package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	res, err := http.Get("http://goosdsfsdfsdfsdfgle.com.br")
	if err != nil {
		log.Fatal("erro ao fazer comunicacao")
	}

	fmt.Println(res.Header)

}
