package database

import (
	"api-go-crud/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Variável global para manter a conexão ativa
var DB *sql.DB

// Abre aconexão com o banco de dados
func ConectarDb() error {
	var erro error

	DB, erro = sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return erro
	}

	if erro = DB.Ping(); erro != nil {
		DB.Close()
		return erro
	}
	return nil
}
