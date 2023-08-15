package main

import (
	"api/src/config"
	router "api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	//fmt.Println(config.Porta)
	//fmt.Println(config.StringConexaoBanco)
	fmt.Printf("Escutando na porta %d", config.Porta)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
