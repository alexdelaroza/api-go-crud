package database

import (
	"fmt"

	"api-go-crud/src/models"
)

// Log
func Log_Inserir(input_log models.Log_input) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	query := `INSERT INTO log (
                          descricao
						, cod_recurso
						, criado_por
                 ) VALUES (?, ?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		msg = fmt.Sprintf("Erro ao preparar a query: %s", err.Error())
		return msg, err
	}

	res, err := stmt.Exec(input_log.Descricao, input_log.Codigo_recurso, input_log.Criado_por)
	if err != nil {
		msg = fmt.Sprintf("Erro ao executar a insercao: %s", err.Error())
		return msg, err
	}

	linhas, err := res.RowsAffected()
	if err != nil {
		msg = fmt.Sprintf("Erro ao validar linhas afetadas: %s", err.Error())
		return msg, err
	}

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) inserida(s).", linhas)
	return msg, nil
}

func Log_Consultar_Codigo(codigo_recurso string) ([]models.Log_output, bool, error, string) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return nil, false, err, msg
	}
	defer db.Close()

	query := "SELECT codigo, descricao, cod_recurso, criado_por, data_criacao_atu FROM log WHERE cod_recurso = ?"

	rows, err := db.Query(query, codigo_recurso)
	if err != nil {
		return nil, false, err, err.Error()
	}
	defer rows.Close()

	var logs []models.Log_output
	for rows.Next() {
		var l models.Log_output
		err := rows.Scan(&l.Codigo, &l.Descricao, &l.Codigo_recurso, &l.Criado_por, &l.Data_criacao_atu)
		if err != nil {
			return nil, false, err, err.Error()
		}
		logs = append(logs, l)
	}

	if err = rows.Err(); err != nil {
		return nil, false, err, err.Error()
	}

	// Se a lista for 0, não achou registros
	if len(logs) == 0 {
		msg := fmt.Sprintf("Nenhum log encontrado para o recurso: %s", codigo_recurso)
		return logs, false, nil, msg
	}

	msg = "Sucesso - Consulta efetuada por codigo"
	return logs, true, nil, msg
}

func Log_Consultar_Data(data_log string) ([]models.Log_output, bool, error, string) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return nil, false, err, msg
	}
	defer db.Close()

	query := "SELECT codigo, descricao, cod_recurso, criado_por, data_criacao_atu FROM log WHERE DATE(data_criacao_atu) = ?"

	rows, err := db.Query(query, data_log)
	if err != nil {
		return nil, false, err, err.Error()
	}
	defer rows.Close()

	var logs []models.Log_output
	for rows.Next() {
		var l models.Log_output
		err := rows.Scan(&l.Codigo, &l.Descricao, &l.Codigo_recurso, &l.Criado_por, &l.Data_criacao_atu)
		if err != nil {
			return nil, false, err, err.Error()
		}
		logs = append(logs, l)
	}

	if err = rows.Err(); err != nil {
		return nil, false, err, err.Error()
	}

	// Se a lista for 0, não achou registros
	if len(logs) == 0 {
		msg := fmt.Sprintf("Nenhum registro encontrado para a data: %s ", data_log)
		return logs, false, nil, msg
	}

	msg = "Sucesso - Consulta efetuada"
	return logs, true, nil, msg
}

func Log_Consultar_Codigo_Data(codigo_recurso, data_log string) ([]models.Log_output, bool, error, string) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return nil, false, err, msg
	}
	defer db.Close()

	query := "SELECT codigo, descricao, cod_recurso, criado_por, data_criacao_atu FROM log WHERE cod_recurso = ? and DATE(data_criacao_atu) = ?"

	rows, err := db.Query(query, codigo_recurso, data_log)
	if err != nil {
		return nil, false, err, err.Error()
	}
	defer rows.Close()

	var logs []models.Log_output
	for rows.Next() {
		var l models.Log_output
		err := rows.Scan(&l.Codigo, &l.Descricao, &l.Codigo_recurso, &l.Criado_por, &l.Data_criacao_atu)
		if err != nil {
			return nil, false, err, err.Error()
		}
		logs = append(logs, l)
	}

	if err = rows.Err(); err != nil {
		return nil, false, err, err.Error()
	}

	// Se a lista for 0, não achou registros
	if len(logs) == 0 {
		msg := fmt.Sprintf("Nenhum registro encontrado para a data: %s ", data_log)
		return logs, false, nil, msg
	}

	msg = "Sucesso - Consulta efetuada"
	return logs, true, nil, msg
}
