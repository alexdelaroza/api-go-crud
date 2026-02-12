package schemas

import (
	"time"
)

type Usuario struct {
	cod_usuario         int
	nome_usuario        string
	login_usuario       string
	senha_usuario       string
	email_usuario       string
	tipo_usuario        string
	ata_ult_atu_usuario time.Time
}
