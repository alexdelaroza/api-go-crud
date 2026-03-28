package validation

import (
	"api-go-crud/src/models"
	"fmt"
	"strings"
	"time"
)

// Validacoes - Gerais
func ValidarData(dateStr1, dateStr2 string) (bool, string) {
	layout := "2006-01-02"

	t1, err1 := time.Parse(layout, dateStr1)
	if err1 != nil {
		return false, "ERRO: Formato da Data-Inicio inválido (use AAAA-MM-DD)"
	}

	t2, err2 := time.Parse(layout, dateStr2)
	if err2 != nil {
		return false, "ERRO: Formato da Data-Fim inválido (use AAAA-MM-DD)"
	}

	// Data de início for MAIOR que Data de fim, é INVÁLIDO
	if t1.After(t2) {
		msg := fmt.Sprintf("Data de Início (%s) não pode ser maior que a Data de Fim (%s)", dateStr1, dateStr2)
		return false, msg // Retorna FALSE para bloquear o processo
	}

	if t1.Equal(t2) {
		return true, "As datas são iguais (Válidas)"
	}

	return true, "As datas são válidas"
}

func ValidarId(id string) (bool, string) {
	// 1. Limpa os espaços primeiro
	idLimpo := strings.TrimSpace(id)

	// 2. Verifica se sobrou algo após a limpeza
	if idLimpo == "" {
		return false, "O campo 'id' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

// Validacoes - Servicos
func ValidarInputServicos(servicos *models.Servico_input) (bool, string) {
	// Limpa os espaços primeiro
	servicos.Descricao = strings.TrimSpace(servicos.Descricao)

	if servicos.Descricao == "" {
		return false, "O campo 'descricao' é obrigatório e deve ser preenchido!"
	}

	if servicos.Valor <= 0 {
		return false, "O campo 'valor' deve ser maior que zero!"
	}

	return true, ""
}

// Validacoes - Usuarios
func ValidaLoginUsuarios(usuario *models.Usuario_login) (bool, string) {
	// Limpa os espaços e converte o e-mail para minúsculas
	usuario.Login = strings.TrimSpace(usuario.Login)
	//usuario.Email = strings.TrimSpace(strings.ToLower(usuario.Email)) // <-- Conversão para minúsculas
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)

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

func ValidarInputUsuarios(usuario *models.Usuario_input) (bool, string) {
	// Limpa os campos primeiro
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Login = strings.TrimSpace(usuario.Login)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Tipo = strings.TrimSpace(usuario.Tipo)

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
