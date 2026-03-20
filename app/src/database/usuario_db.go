package database

import (
	"fmt"

	"api-go-crud/src/models"
)

// Usuarios
func Usuario_Inserir(novo_usuario models.Usuario_input) (int, string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return 0, msg, err
	}
	defer db.Close()

	query := `
		INSERT INTO usuarios ( 
	                nome, 
					login,
					senha, 
					email, 
					tipo
           ) VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		msg = fmt.Sprintf("Erro ao preparar a query: %s", err.Error())
		return 0, msg, err
	}

	res, err := stmt.Exec(novo_usuario.Nome, novo_usuario.Login, novo_usuario.Senha, novo_usuario.Email, novo_usuario.Tipo)
	if err != nil {
		msg = fmt.Sprintf("Erro ao executar a insercao: %s", err.Error())
		return 0, msg, err
	}

	id, err := res.LastInsertId()
	//fmt.Println(id)

	linhas, err := res.RowsAffected()
	if err != nil {
		msg = fmt.Sprintf("Erro ao validar linhas afetadas: %s", err.Error())
		return 0, msg, err
	}

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) inserida(s).", linhas)
	return int(id), msg, err
}

func Usuario_Atualizar(codigo string, altera_usuario models.Usuario_input) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	query := `
		update usuarios 
	       set   nome   = ? 
		     ,   login  = ?
			 ,   senha  = ? 
			 ,   email  = ? 
			 ,   tipo   = ?
         where   codigo = ?
	`

	stmt, _ := db.Prepare(query)

	res, err := stmt.Exec(altera_usuario.Nome, altera_usuario.Login, altera_usuario.Senha, altera_usuario.Email, altera_usuario.Tipo, codigo)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada para ser retornada
	msg = fmt.Sprintf("Sucesso! %d linha(s) afetada(s).", linhas)
	return msg, nil
}

func Usuario_Deletar(codigo_usuario string) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	query := `
		delete from usuarios 
		 where codigo = ?
	`

	stmt, _ := db.Prepare(query)

	res, _ := stmt.Exec(codigo_usuario)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) deletada(s).", linhas)
	return msg, nil
}

func Usuario_Consultar() ([]models.Usuario_read, error, string) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return nil, err, msg
	}
	defer db.Close()

	var usuarios []models.Usuario_read
	query := `
		SELECT codigo, nome, login, email, data_criacao_atu 
		FROM usuarios
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var u models.Usuario_read
		err := rows.Scan(&u.Codigo, &u.Nome, &u.Login, &u.Email, &u.Data_criacao_atu)
		if err != nil {
			return nil, err, err.Error()
		}
		usuarios = append(usuarios, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err, err.Error()
	}

	msg = "Sucesso - Consulta efetuada"
	return usuarios, nil, msg
}

func Usuario_Consultar_Codigo(codigo_usuario string) (models.Usuario_read, bool, error, string) {
	var msg string
	var usuario models.Usuario_read

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return usuario, false, err, msg
	}
	defer db.Close()

	query := "select codigo, nome, login, email, data_criacao_atu from usuarios where codigo = ?"

	rows, err := db.Query(query, codigo_usuario)
	if err != nil {
		return usuario, false, err, err.Error()
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o código: %s", codigo_usuario)
		return usuario, false, nil, msg
	}

	err = rows.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Login, &usuario.Email, &usuario.Data_criacao_atu)
	if err != nil {
		// Erro real
		return usuario, false, err, err.Error()
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Sucesso - Usuario %s encontrado com sucesso", usuario.Codigo)
	return usuario, true, nil, msg
}

func Usuario_Consultar_Email(email_usuario string) (bool, string, error) {
	var msg string
	var total int

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return false, msg, err
	}
	defer db.Close()

	query := `
	    SELECT COUNT(*)
		  FROM usuarios
		 WHERE email = ? 
	`

	rows, err := db.Query(query, email_usuario)
	if err != nil {
		return false, err.Error(), err
	}
	defer rows.Close()

	rows.Next()

	err = rows.Scan(&total)
	if err != nil {
		msg = "Erro ao validar e-mail"
		return false, msg, err // Erro real
	}

	if total == 0 {
		msg = fmt.Sprintf("Nenhum registro encontrado para o email: %s ", email_usuario)
		return false, msg, nil
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Email: %s cadastrado", email_usuario)
	return true, msg, nil
}

func Usuario_Consultar_Login(login_usuario string) (bool, string, error) {
	var msg string
	var total int

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return false, msg, err
	}
	defer db.Close()

	query := `
	    SELECT COUNT(*)
		  FROM usuarios
		 WHERE login = ? 
	`

	rows, err := db.Query(query, login_usuario)
	if err != nil {
		return false, err.Error(), err
	}
	defer rows.Close()

	rows.Next()

	err = rows.Scan(&total)
	if err != nil {
		msg = "Erro ao validar login"
		return false, msg, err // Erro real
	}

	if total == 0 {
		msg = fmt.Sprintf("Nenhum registro encontrado para o email: %s ", login_usuario)
		return false, msg, nil
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Login: %s cadastrado", login_usuario)
	return true, msg, nil
}
