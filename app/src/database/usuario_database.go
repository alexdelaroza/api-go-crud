package database

import (
	"fmt"

	"api-go-crud/src/models"
)

// Usuarios
func Usuario_Inserir(novo_usuario models.Usuario) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
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

	res, err := stmt.Exec(novo_usuario.Codigo, novo_usuario.Nome, novo_usuario.Login, novo_usuario.Senha, novo_usuario.Email, novo_usuario.Tipo)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) inserida(s).", linhas)
	return msg, nil
}

func Usuario_Atualizar(altera_usuario models.Usuario) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
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

	res, err := stmt.Exec(altera_usuario.Nome, altera_usuario.Login, altera_usuario.Senha, altera_usuario.Email, altera_usuario.Tipo, altera_usuario.Codigo)

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

	query := `delete from usuarios where cod_usuario = ?`

	stmt, _ := db.Prepare(query)

	res, _ := stmt.Exec(codigo_usuario)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) deletada(s).", linhas)
	return msg, nil
}

func Usuario_Consultar() ([]models.Usuario, error, string) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return nil, err, msg
	}
	defer db.Close()

	var usuarios []models.Usuario
	query := `SELECT cod_usuario, nome_usuario, login_usuario, senha_usuario, email_usuario, tipo_usuario, data_ult_atu_usuario FROM usuarios`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(&u.Codigo, &u.Nome, &u.Login, &u.Senha, &u.Email, &u.Tipo, &u.Data_ult_atu)
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

func Usuario_Consultar_Codigo(codigo_usuario string) (models.Usuario, bool, error, string) {
	var msg string
	var usuario models.Usuario

	db, err := Conectar()
	if err != nil {
		fmt.Println("erro 1")
		return usuario, false, err, err.Error()
	}
	defer db.Close()

	query := "select cod_usuario, nome_usuario, login_usuario, senha_usuario, email_usuario, tipo_usuario, data_ult_atu_usuario from usuarios where cod_usuario = ?"

	rows, err := db.Query(query, codigo_usuario)
	if err != nil {
		return usuario, false, err, err.Error()
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o código: %s", codigo_usuario)
		return usuario, false, nil, msg
	}

	err = rows.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Login, &usuario.Senha, &usuario.Email, &usuario.Tipo, &usuario.Data_ult_atu)
	if err != nil {
		// Erro real
		fmt.Println("erro 2")
		return usuario, false, err, err.Error()
	}

	fmt.Println("achou 3")
	// Sucesso - Encontrou
	msg = "Sucesso - Consulta efetuada"
	return usuario, true, nil, msg
}

// Servicos
func Servico_Consultar() ([]models.Servico, error, string) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return nil, err, msg
	}
	defer db.Close()

	query := `SELECT cod_servico, descricao_servico, valor_servico, data_ult_atu_servico FROM servico`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()

	var servicos []models.Servico

	for rows.Next() {
		var s models.Servico
		err := rows.Scan(&s.Codigo, &s.Descricao, &s.Valor, &s.Data_ult_atu)
		if err != nil {
			return nil, err, err.Error()
		}
		servicos = append(servicos, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err, err.Error()
	}

	msg = "Sucesso - Consulta efetuada"
	return servicos, nil, msg
}

func Servico_Consultar_Codigo(codigo_servico string) (models.Servico, bool, error, string) {
	var msg string

	var servico models.Servico
	db, err := Conectar()
	if err != nil {
		fmt.Println("erro 1")
		return servico, false, err, err.Error()
	}
	defer db.Close()

	query := "SELECT cod_servico, descricao_servico, valor_servico, data_ult_atu_servico FROM servico WHERE cod_servico = ?"

	rows, err := db.Query(query, codigo_servico)
	if err != nil {
		return servico, false, err, err.Error()
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o código: %s ", codigo_servico)
		return servico, false, nil, msg
	}

	err = rows.Scan(&servico.Codigo, &servico.Descricao, &servico.Valor, &servico.Data_ult_atu)

	if err != nil {
		fmt.Println("entrou 2")
		return servico, false, err, err.Error() // Erro real
	}

	fmt.Println("entrou 3")
	// Sucesso - Encontrou
	msg = "Sucesso - Consulta efetuada"
	return servico, true, nil, msg
}
