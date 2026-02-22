package models

import (
	"time"
)

type Log struct {
	Codigo         string
	Descricao      string
	Codigo_recurso int
	Criado_por     string
	Data_ult_atu   time.Time
}
