package validation

import (
	"api-go-crud/src/models"
	"fmt"
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
	if id == "" {
		return false, "O campo 'id' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

// Validacoes - Servicos
func ValidarInputServicos(servico models.Servico_input) (bool, string) {
	if servico.Descricao == "" {
		return false, "O campo 'descricao' é obrigatório e deve ser preenchido!"
	}

	if servico.Valor == 0 {
		return false, "O campo 'valor' é obrigatório e deve ser preenchido!"
	}
	return true, ""
}

// Validacoes - Usuarios
func ValidaLoginUsuario(usuario models.Usuario_login) (bool, string) {

	if usuario.Login == "" && usuario.Email == "" {
		return false, "O campo 'email' ou 'login' é obrigatório e deve ser preenchido!"
	}

	if usuario.Senha == "" {
		return false, "O campo 'senha' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

func ValidarInputUsuario(usuario models.Usuario_input) (bool, string) {
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

	if usuario.Tipo == "" {
		return false, "O campo 'tipo' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}
