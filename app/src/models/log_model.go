package models

import (
	"time"
)

type Log_output struct {
	Codigo         string
	Descricao      string
	Codigo_recurso string
	Criado_por     string
	Data_ult_atu   time.Time
}

type Log_input struct {
	Descricao      string
	Codigo_recurso string
	Criado_por     string
}
