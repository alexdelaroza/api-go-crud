package database

import (
	"fmt"

	"api-go-crud/src/models"
)

// Servicos
func Servico_Inserir(novo_servico models.Servico) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	query := `INSERT INTO servico (
                          cod_servico
                        , descricao_servico
						, valor_servico
                 ) VALUES (?, ?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		msg = fmt.Sprintf("Erro ao preparar a query: %s", err.Error())
		return msg, err
	}

	res, err := stmt.Exec(novo_servico.Codigo, novo_servico.Descricao, novo_servico.Valor)
	if err != nil {
		msg = fmt.Sprintf("Erro ao executar a insercao: %s", err.Error())
		return msg, err
	}

	id, err := res.LastInsertId()
	fmt.Println(id)

	linhas, err := res.RowsAffected()
	if err != nil {
		msg = fmt.Sprintf("Erro ao validar linhas afetadas: %s", err.Error())
		return msg, err
	}

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) inserida(s).", linhas)
	return msg, nil
}

func Servico_Atualizar(altera_servico models.Servico) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	query := `update  servico 
	            set   descricao_servico = ? 
				  ,   valor_servico     = ?
                where cod_servico       = ?`

	stmt, _ := db.Prepare(query)

	res, err := stmt.Exec(altera_servico.Descricao, altera_servico.Valor, altera_servico.Codigo)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada para ser retornada
	msg = fmt.Sprintf("Sucesso! %d linha(s) afetada(s).", linhas)
	return msg, nil
}

func Servico_Deletar(codigo_servico string) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	query := `delete from servico where cod_servico = ?`

	stmt, _ := db.Prepare(query)

	res, _ := stmt.Exec(codigo_servico)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) deletada(s).", linhas)
	return msg, nil
}

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
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return servico, false, err, msg
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
		return servico, false, err, err.Error() // Erro real
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Sucesso - Servico %s encontrado com sucesso", servico.Codigo)
	return servico, true, nil, msg
}
