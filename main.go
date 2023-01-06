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
	router.HandleFunc("/usuarios", usuario.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", usuario.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", usuario.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", usuario.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta: 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
