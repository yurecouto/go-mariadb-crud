package usuario

import (
	"database/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha converter id"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Falha conectar no banco"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Falha ao buscar usuario"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Falha ao escanear usuario"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario nao encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Falha converter usuario para json"))
		return
	}

}
