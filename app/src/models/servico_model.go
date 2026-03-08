package models

import (
	"time"
)

type Servico_output struct {
	Codigo           string
	Descricao        string
	Valor            float64
	Data_criacao_atu time.Time
}

type Servico_input struct {
	Descricao string
	Valor     float64
}
