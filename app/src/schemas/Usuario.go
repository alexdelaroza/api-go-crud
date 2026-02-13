package schemas

import (
	"time"
)

type Usuario struct {
	Codigo      int
	Nome        string
	Login       string
	Senha       string
	Email       string
	Tipo        string
	Data_ult_atu time.Time
}
