package validation

import (
	"api-go-crud/src/models"
	"strings"
)

// Validacoes - Servicos
func ValidarInputServicos(servicos models.Servico_input) (bool, string) {
	LimparInputServicos(&servicos)

	if servicos.Descricao == "" {
		return false, "O campo 'descricao' é obrigatório e deve ser preenchido!"
	}

	if servicos.Valor <= 0 {
		return false, "O campo 'valor' deve ser maior que zero!"
	}

	return true, ""
}

func LimparInputServicos(servicos *models.Servico_input) {
	// Limpa os espaços dos campos
	servicos.Descricao = strings.TrimSpace(servicos.Descricao)

}
