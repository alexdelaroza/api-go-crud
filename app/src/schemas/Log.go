package schemas

import (
	"time"
)

type Log struct {
	Codigo         int
	Descricao      string
	Codigo_recurso int
	Tipo           string
	Data_ult_atu_  time.Time
}
