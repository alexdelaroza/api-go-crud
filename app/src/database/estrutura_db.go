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

func Insert_regs_usuarios() {
	// Garante o uso do banco correto
	exec(DB, "USE crud_db")

	// Inserção em massa (Multi-insert) para ser mais eficiente
	exec(DB, `
		INSERT INTO usuarios 
			(codigo, nome, login, senha, email, tipo, data_criacao_atu)
		VALUES
			(1,  'marcela nobrega',   'marcelanobrega',  '123', 'marcelanobrega@gmail.com',    'admin',    '2026-03-08 00:02:43'),
			(4,  'alex delaroza',     'adelaroza',       '123', 'alex.delaroza@gmail.com',     'user',     '2026-05-07 01:38:50'),
			(5,  'alisson delaroza',  'alissondelaroza', '123', 'alisson.delaroza@gmail.com',  'admin',    '2026-03-08 03:20:18'),
			(10, 'neuza delaroza',    'neuza.delaroza',  '123', 'neuza.delaroza@gmail.com.br', 'admin',    '2026-03-17 02:39:53'),
			(11, 'romeu',             'romeu',           '123', 'romeu@gmail.com.br',          'vendedor', '2026-04-12 00:07:29')
	`)

}

func Insert_regs_servicos() {
	// Garante o uso do banco correto
	exec(DB, "USE crud_db")

	// Inserção em massa (Multi-insert) para ser mais eficiente
	exec(DB, `
		INSERT INTO servico (codigo, descricao, valor, data_criacao_atu)
		VALUES
			(1, 'servico_001', 101.00, '2026-03-08 00:03:11'),
			(2, 'servico_002', 201.00, '2026-03-08 00:32:55'),
			(3, 'servico_003', 301.00, '2026-03-08 01:42:55'),
			(4, 'servico_004', 401.00, '2026-03-08 20:01:22'),
			(5, 'servico_005', 501.00, '2026-03-08 20:02:05'),
			(7, 'servico 006', 300.01, '2026-04-12 20:07:59'),
			(8, 'serviço 007', 3255.00, '2026-05-10 00:59:34')
	`)
}
