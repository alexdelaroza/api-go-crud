package database

import (
	"fmt"
	"log"

	"api-go-crud/src/models"
)

func Usuario_Inserir(novo_usuario models.Usuario) {
	db, err := Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	query := `INSERT INTO usuarios (
                          cod_usuario, 
						  nome_usuario, 
						  login_usuario,
						  senha_usuario, 
						  email_usuario, 
						  tipo_usuario
                 ) VALUES (?, ?, ?, ?, ?, ?)`

	stmt, _ := db.Prepare(query)

	//res, _ := stmt.Exec(1, // cod_usuario
	//	"Maria",           // nome_usuario
	//	"mariasilva",      // login_usuario
	//	"senha123",        // senha_usuario
	//	"maria@email.com", // email_usuario
	//	"admin",           // tipo_usuario
	//)

	// Passamos os campos da struct para o Exec
	res, err := stmt.Exec(novo_usuario.Codigo, novo_usuario.Nome, novo_usuario.Login, novo_usuario.Senha, novo_usuario.Email, novo_usuario.Tipo)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Sucesso! %d linha(s) afetada(s).\n", linhas)
}

func Usuario_Atualizar() {
	db, err := Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	query := `update usuarios 
	            set   cod_usuario   = ? 
				  ,   nome_usuario  = ? 
				  ,   login_usuario = ?
				  ,   senha_usuario = ? 
				  ,   email_usuario = ? 
				  ,   tipo_usuario  = ?
                where cod_usuario   = ?`

	stmt, _ := db.Prepare(query)

	res, _ := stmt.Exec(1, // cod_usuario
		"Antonio",           // nome_usuario
		"antoniocarlos",     // login_usuario
		"senha123",          // senha_usuario
		"antonio@email.com", // email_usuario
		"admin",             // tipo_usuario
		1,                   // cod_usuario
	)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Sucesso! %d linha(s) afetada(s).\n", linhas)
}

func Usuario_Deletar() {
	db, err := Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	query := `delete from usuarios where cod_usuario = ?`

	stmt, _ := db.Prepare(query)

	res, _ := stmt.Exec(1) // cod_usuario

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Sucesso! %d linha(s) afetada(s).\n", linhas)
}

func Usuario_Consultar() {
	db, err := Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios where cod_usuario = ?", 1)
	defer db.Close()

	for rows.Next() {
		var usuario models.Usuario
		rows.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Login, &usuario.Senha, &usuario.Email, &usuario.Tipo, &usuario.Data_ult_atu)
		fmt.Printf("Usuário: %d - %s (Atualizado em: %s)\n",
			usuario.Codigo, usuario.Nome, usuario.Data_ult_atu.Format("02/01/2006 15:04:05"))
	}
}
