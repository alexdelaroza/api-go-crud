package validation

import (
	"api-go-crud/src/models"
	"strings"
)

// Validacoes - Usuarios
func ValidaLoginUsuarios(usuario models.Usuario_login) (bool, string) {
	LimparLoginUsuarios(&usuario)

	if usuario.Login == "" && usuario.Email == "" {
		return false, "O campo 'email' ou 'login' é obrigatório e deve ser preenchido!"
	}

	if usuario.Email != "" && !strings.Contains(usuario.Email, "@") {
		return false, "O e-mail informado é inválido!"
	}

	if usuario.Senha == "" {
		return false, "O campo 'senha' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

func LimparLoginUsuarios(usuario *models.Usuario_login) {
	// Limpa os espaços dos campos
	//usuario.Email = strings.TrimSpace(strings.ToLower(usuario.Email)) // <-- Conversão para minúsculas
	usuario.Login = strings.TrimSpace(usuario.Login)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)

}

func ValidarInputUsuarios(usuario models.Usuario_input) (bool, string) {
	LimparInputUsuarios(&usuario)

	if usuario.Nome == "" {
		return false, "O campo 'nome' é obrigatório e deve ser preenchido!"
	}

	if usuario.Login == "" {
		return false, "O campo 'login' é obrigatório e deve ser preenchido!"
	}

	if usuario.Senha == "" {
		return false, "O campo 'senha' é obrigatório e deve ser preenchido!"
	}

	if usuario.Email == "" {
		return false, "O campo 'email' é obrigatório e deve ser preenchido!"
	}

	if !strings.Contains(usuario.Email, "@") {
		return false, "O e-mail informado é inválido!"
	}

	if usuario.Tipo == "" {
		return false, "O campo 'tipo' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

func LimparInputUsuarios(usuario *models.Usuario_input) {
	// Limpa os espaços dos campos
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Login = strings.TrimSpace(usuario.Login)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Tipo = strings.TrimSpace(usuario.Tipo)
}
