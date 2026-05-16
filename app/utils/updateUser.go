package utils

import "fmt"

// UpdateUser atualiza os dados de um usuário no banco pelo e-mail atual
func UpdateUser(currentEmail, newUsername, newEmail, newBornDate, newPassword string) error {
	db, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("erro ao criptografar senha: %w", err)
	}

	result, err := db.Exec(
		`UPDATE users SET username=$1, email=$2, born_date=$3, password=$4 WHERE email=$5`,
		newUsername, newEmail, newBornDate, hashedPassword, currentEmail,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("nenhum usuário encontrado com o e-mail informado")
	}

	return nil
}
