package utils

import (
	"fmt"
	"time"
)

// User representa um usuário do sistema de saúde
type User struct {
	ID        int
	Username  string
	Email     string
	BornDate  string
	Password  string
	CreatedAt time.Time
}

// GetUserByEmail busca um usuário pelo e-mail no banco de dados
func GetUserByEmail(email string) (*User, error) {
	db, err := ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &User{}
	row := db.QueryRow(
		"SELECT id, username, email, born_date, password, created_at FROM users WHERE email = $1",
		email,
	)

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.BornDate, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado: %w", err)
	}

	return user, nil
}
