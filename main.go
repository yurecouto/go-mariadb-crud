package main

import (
	"database/modules/usuario"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", usuario.CreateUser).Methods(http.MethodPost)

	fmt.Println("Escutando na porta: 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
