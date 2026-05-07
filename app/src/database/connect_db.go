package database

import (
	"api-go-crud/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Abre aconexão com o banco de dados
func ConectarDb() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
