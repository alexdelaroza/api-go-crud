package models

import (
	"time"
)

type Usuario_input struct {
	Nome         string
	Login        string
	Senha        string
	Email        string
	Tipo         string
}

type Usuario_output struct {
	Codigo       string
	Nome         string
	Login        string
	Senha        string
	Email        string
	Tipo         string
	Data_criacao_atu time.Time
}

type Usuario_read struct {
	Codigo       string
	Nome         string
	Login        string
	Email        string
	Data_criacao_atu time.Time
}

type Usuario_login struct {
	Login        string
	Senha        string
	Email        string
}