package schemas

import (
	"time"
)

type Log struct {
	cod_log          int
	descricao_log    string
	cod_recurso      int
	tipo_log         string
	data_ult_atu_log time.Time
}
