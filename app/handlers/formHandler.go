package handlers

import (
	"fmt"
	"log"
	"net/http"

	"servidorHTTP/app/utils"
)

// FormHandler processa a criação de novas contas de usuário
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	bornDate := r.FormValue("born_date")
	password := r.FormValue("password")

	if err := utils.ValidateUser(username, email, password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Erro interno ao processar senha", http.StatusInternalServerError)
		return
	}

	db, err := utils.ConnectToDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users (username, email, born_date, password) VALUES ($1, $2, $3, $4)",
		username, email, bornDate, hashedPassword,
	)
	if err != nil {
		log.Printf("Erro ao criar usuário: %v", err)
		http.Error(w, "E-mail já cadastrado ou erro ao criar conta", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<script>alert("✅ Conta criada com sucesso!"); window.location.href="/";</script>`)
}
