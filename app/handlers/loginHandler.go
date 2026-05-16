package handlers

import (
	"fmt"
	"net/http"

	"servidorHTTP/app/utils"
)

// LoginHandler valida as credenciais e exibe o perfil do usuário
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := utils.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPassword(password, user.Password) {
		http.Error(w, "Senha incorreta", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <title>Perfil - Sistema de Saúde</title>
  <link rel="stylesheet" href="/styles/login.style.css">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link href="https://fonts.googleapis.com/css2?family=DM+Serif+Display&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
</head>
<body>
  <div class="profile-container">
    <div class="profile-card">
      <div class="profile-badge">✅</div>
      <h1>Bem-vindo, %s!</h1>
      <div class="profile-info">
        <div class="info-row"><span class="label">E-mail</span><span class="value">%s</span></div>
        <div class="info-row"><span class="label">Data de Nascimento</span><span class="value">%s</span></div>
        <div class="info-row"><span class="label">Membro desde</span><span class="value">%s</span></div>
      </div>
      <a href="/" class="btn-back">← Voltar ao início</a>
    </div>
  </div>
</body>
</html>
`, user.Username, user.Email, user.BornDate, user.CreatedAt.Format("02/01/2006"))
}
