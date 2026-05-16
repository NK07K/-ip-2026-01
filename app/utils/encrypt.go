package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword gera um hash bcrypt para a senha fornecida
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar hash: %w", err)
	}
	return string(bytes), nil
}

// CheckPassword compara uma senha em texto com seu hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
