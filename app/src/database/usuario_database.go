package database

import (
	"database/sql"
	"fmt"
	"log"

	"api-go-crud/src/models"
)

func Usuario_Inserir(novo_usuario models.Usuario) (string, error) {
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
	// fmt.Sprintf cria a string formatada para ser retornada
	mensagem := fmt.Sprintf("Sucesso! %d linha(s) afetada(s).", linhas)

	return mensagem, nil
}

func Usuario_Atualizar(altera_usuario models.Usuario) (string, error) {
	db, err := Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	query := `update usuarios 
	            set   nome_usuario  = ? 
				  ,   login_usuario = ?
				  ,   senha_usuario = ? 
				  ,   email_usuario = ? 
				  ,   tipo_usuario  = ?
                where cod_usuario   = ?`

	stmt, _ := db.Prepare(query)

	//res, _ := stmt.Exec(1, // cod_usuario
	//	"Antonio",           // nome_usuario
	//	"antoniocarlos",     // login_usuario
	//	"senha123",          // senha_usuario
	//	"antonio@email.com", // email_usuario
	//	"admin",             // tipo_usuario
	//)

	// Passamos os campos da struct para o Exec
	res, err := stmt.Exec(altera_usuario.Nome, altera_usuario.Login, altera_usuario.Senha, altera_usuario.Email, altera_usuario.Tipo, altera_usuario.Codigo)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	// fmt.Sprintf cria a string formatada para ser retornada
	mensagem := fmt.Sprintf("Sucesso! %d linha(s) afetada(s).", linhas)

	return mensagem, nil
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

	for rows.Next() {
		var usuario models.Usuario
		rows.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Login, &usuario.Senha, &usuario.Email, &usuario.Tipo, &usuario.Data_ult_atu)
		fmt.Printf("Usuário: %d - %s (Atualizado em: %s)\n",
			usuario.Codigo, usuario.Nome, usuario.Data_ult_atu.Format("02/01/2006 15:04:05"))
	}
}

func Usuario_Consultar_Codigo(codigo_usuario string) (models.Usuario, bool, error) {
	var usuario models.Usuario
	db, err := Conectar()
	if err != nil {
		return usuario, false, err
	}
	defer db.Close()

	rows, err := db.Query("select * from usuarios where cod_usuario = ?", codigo_usuario)

	for rows.Next() {
		var usuario models.Usuario
		rows.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Login, &usuario.Senha, &usuario.Email, &usuario.Tipo, &usuario.Data_ult_atu)
		//fmt.Printf("Usuário: %d - %s (Atualizado em: %s)\n",
		//	usuario.Codigo, usuario.Nome, usuario.Data_ult_atu.Format("02/01/2006 15:04:05"))
	}

	if err != nil {
		if err == sql.ErrNoRows {
			// Não encontrou, mas não é um "erro de sistema". Retorna falso.
			return usuario, false, nil
		}
		return usuario, false, err // Erro real (conexão, sintaxe, etc)
	}

	return usuario, true, nil // Encontrou com sucesso
}
