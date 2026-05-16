package utils

import "fmt"

// DeleteUser remove um usuário do banco pelo e-mail após validar a senha
func DeleteUser(email, password string) error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("usuário não encontrado")
	}

	if !CheckPassword(password, user.Password) {
		return fmt.Errorf("senha incorreta")
	}

	db, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE email = $1", email)
	if err != nil {
		return fmt.Errorf("erro ao excluir usuário: %w", err)
	}

	return nil
}
