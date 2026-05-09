package database

import (
	"database/sql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func Create_table() {

	//exec(db, "create database if not exists cursogo")
	//exec(db, "use cursogo")

	exec(DB, "create database if not exists crud_db")
	exec(DB, "use crud_db")

	exec(DB, "drop table if exists log")
	exec(DB, "drop table if exists servico")
	exec(DB, "drop table if exists usuarios")

	exec(DB, `create table usuarios (
             codigo INT NOT NULL AUTO_INCREMENT,
             nome VARCHAR(250) NOT NULL,
             login VARCHAR(250) NOT NULL,
             senha VARCHAR(250) NOT NULL,
             email VARCHAR(250) NOT NULL UNIQUE,
             tipo VARCHAR(250) NOT NULL,
             data_criacao_atu TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
             PRIMARY KEY (codigo)
        )`)

	exec(DB, `create table servico (
             codigo INT NOT NULL AUTO_INCREMENT,
             descricao VARCHAR(250) NOT NULL,
             valor DECIMAL(10, 2) NOT NULL,
             data_criacao_atu TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
             CONSTRAINT chk_valor_positivo CHECK (valor > 0),
             PRIMARY KEY (codigo)
        )`)

	exec(DB, `create table log (
             codigo INT NOT NULL AUTO_INCREMENT,
             descricao TEXT NOT NULL,
             cod_recurso VARCHAR(36) NOT NULL,
             criado_por INT NOT NULL,
             data_criacao_atu TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
             PRIMARY KEY (codigo),
             foreign key(criado_por) references usuarios(codigo)
        )`)

}
