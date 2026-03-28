package database

import (
	"api-go-crud/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Abre aconexão com o banco de dados
func Conectar() (*sql.DB, error) {
	//stringConexaoBanco := "root:root@/crud_db?parseTime=True"

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
