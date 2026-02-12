package schemas

import (
	"time"
)

type Servico struct {
	cod_servico         int
	descricao_servico   string
	valor_servico       float64
	data_ult_atu_servic time.Time
}
