package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Abre aconexão com o banco de dados
func Conectar() (*sql.DB, error) {
	//stringConexao := "root:root@/cursogo?parseTime=True"
	stringConexao := "root:root@/crud_db?parseTime=True"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
