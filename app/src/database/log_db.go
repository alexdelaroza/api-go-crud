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

func Log_Consultar_Codigo(codigo_recurso string) (models.Log_output, bool, error, string) {
	var msg string

	var log models.Log_output
	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return log, false, err, msg
	}
	defer db.Close()

	log.Codigo_recurso = codigo_recurso
	query := "SELECT codigo, descricao, criado_por, data_criacao_atu FROM log WHERE cod_recurso = ?"

	rows, err := db.Query(query, codigo_recurso)
	if err != nil {
		return log, false, err, err.Error()
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o código de recurso: %s ", codigo_recurso)
		return log, false, nil, msg
	}

	err = rows.Scan(&log.Codigo, &log.Descricao, &log.Criado_por, &log.Data_criacao_atu, &log.Codigo_recurso)
	if err != nil {
		return log, false, err, err.Error() // Erro real
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Sucesso - Log do recurso %s encontrado com sucesso", codigo_recurso)
	return log, true, nil, msg
}
