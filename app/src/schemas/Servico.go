package schemas

import (
	"time"
)

type Servico struct {
	Codigo       int
	Descricao    string
	Valor        float64
	Data_ult_atu time.Time
}
