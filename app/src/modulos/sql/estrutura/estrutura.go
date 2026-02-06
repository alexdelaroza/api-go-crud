package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")

	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
             cod_usuario INT,
             nome_usuario VARCHAR(250) NOT NULL,
             login_usuario VARCHAR(250) NOT NULL,
             senha_usuario VARCHAR(250) NOT NULL,
             email_usuario VARCHAR(250) UNIQUE,
             tipo_usuario VARCHAR(250) NOT NULL,
             data_ult_atu_usuario TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
             PRIMARY KEY (cod_usuario)
        )`)

	exec(db, "drop table if exists servico")
	exec(db, `create table servico (
             cod_servico INT,
             descricao_servico VARCHAR(250) NOT NULL,
             valor_servico DECIMAL(10, 2) NOT NULL,
             data_ult_atu_servico TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
             CONSTRAINT chk_valor_positivo CHECK (valor_servico > 0),
             PRIMARY KEY (cod_servico)
        )`)

	exec(db, "drop table if exists log")
	exec(db, `create table log (
             cod_log INT,
             descricao_log TEXT NULL,
             cod_recurso INT NOT NULL,
             tipo_log VARCHAR(250) NOT NULL,
             data_ult_atu_log TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
             PRIMARY KEY (cod_log)
        )`)

}
