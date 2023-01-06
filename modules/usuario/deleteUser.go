package usuario

import (
	"database/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha converter id"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Falha ao deletar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
