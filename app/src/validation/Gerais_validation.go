package validation

import (
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

	switch {
	case t1.After(t2):
		return false, fmt.Sprintf("Data de Início (%s) não pode ser maior que a Data de Fim (%s)", dateStr1, dateStr2)
	case t1.Equal(t2):
		return true, "As datas são iguais (Válidas)"
	default:
		return true, "As datas são válidas"
	}

}

func ValidarId(id string) (bool, string) {
	//Limpa os espaços do id
	idLimpo := strings.TrimSpace(id)

	//Verifica se id esta preenchido
	if idLimpo == "" {
		return false, "O campo 'id' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}
