package authorization

import (
	"api-go-crud/src/models"
)

func CheckRole(usuario models.Usuario_read, tiposPermitidos ...string) (string, error) {
	var msg string

	if usuario.Tipo == "" {
		msg = "O campo 'TIPO' é obrigatório e deve ser preenchido!"
	}

	for _, role := range tiposPermitidos {
		if role == usuario.Tipo {
			msg = "Acesso aos Serviços permitido"

		}
	}

	return msg, nil
}
