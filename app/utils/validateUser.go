package utils

import (
	"fmt"
	"strings"
)

// ValidateUser valida os campos obrigatórios de um usuário
func ValidateUser(username, email, password string) error {
	if strings.TrimSpace(username) == "" {
		return fmt.Errorf("nome de usuário é obrigatório")
	}
	if len(username) < 3 {
		return fmt.Errorf("nome deve ter ao menos 3 caracteres")
	}
	if strings.TrimSpace(email) == "" {
		return fmt.Errorf("e-mail é obrigatório")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("e-mail inválido")
	}
	if len(password) < 6 {
		return fmt.Errorf("senha deve ter ao menos 6 caracteres")
	}
	return nil
}
