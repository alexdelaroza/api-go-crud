package models

import (
	"time"
)

type Usuario struct {
	Codigo       string
	Nome         string
	Login        string
	Senha        string
	Email        string
	Tipo         string
	Data_ult_atu time.Time
}
