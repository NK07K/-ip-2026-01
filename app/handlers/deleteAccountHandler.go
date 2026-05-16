package handlers

import (
	"fmt"
	"net/http"

	"servidorHTTP/app/utils"
)

// DeleteAccountHandler remove a conta do usuário do banco de dados
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if err := utils.DeleteUser(email, password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<script>alert("🗑️ Conta excluída com sucesso."); window.location.href="/";</script>`)
}
