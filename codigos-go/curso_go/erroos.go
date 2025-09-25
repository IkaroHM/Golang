package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://github.com")
	if err != nil {
		log.Println("Falha ao acessar o site, codigo de erro ", err)
		return
	}

		defer res.Body.Close()
	
	log.Println("site acessado com sucesso. Status code:", res.StatusCode)
}