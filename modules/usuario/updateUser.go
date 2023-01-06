package usuario

import (
	"database/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha converter id"))
		return
	}

	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler corpo da requisicao"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(reqBody, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter usuario para struct"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Falha ao atualizar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
