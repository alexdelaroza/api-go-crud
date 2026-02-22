package models

import (
	"time"
)

type Servico struct {
	Codigo       string
	Descricao    string
	Valor        float64
	Data_ult_atu time.Time
}
