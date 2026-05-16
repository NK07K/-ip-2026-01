package handlers

import (
	"fmt"
	"net/http"

	"servidorHTTP/app/utils"
)

// UpdateAccountHandler atualiza os dados do usuário no banco de dados
func UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	currentEmail := r.FormValue("current_email")
	currentPassword := r.FormValue("current_password")
	newUsername := r.FormValue("new_username")
	newEmail := r.FormValue("new_email")
	newBornDate := r.FormValue("new_born_date")
	newPassword := r.FormValue("new_password")

	// Valida credenciais atuais
	user, err := utils.GetUserByEmail(currentEmail)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPassword(currentPassword, user.Password) {
		http.Error(w, "Senha atual incorreta", http.StatusUnauthorized)
		return
	}

	if err := utils.ValidateUser(newUsername, newEmail, newPassword); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.UpdateUser(currentEmail, newUsername, newEmail, newBornDate, newPassword); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<script>alert("✅ Conta atualizada com sucesso!"); window.location.href="/";</script>`)
}
