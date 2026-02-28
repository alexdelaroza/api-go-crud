package database

import (
	"fmt"
	"strconv"

	"api-go-crud/src/models"
)

// Log
func Log_Inserir(novo_log models.Log) (string, error) {
	var msg string

	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return msg, err
	}
	defer db.Close()

	// Busca o maior ID atual
	var ultimoID int
	rows, err := db.Query("SELECT COALESCE(MAX(cod_log), 0) FROM log")
	if err != nil {
		msg = fmt.Sprintf("Erro buscar o ultimo id: %s", err.Error())
		return msg, err
	}
	defer rows.Close()

	rows.Scan(&ultimoID)
	novo_log.Codigo = strconv.Itoa(ultimoID + 1)

	// Busca efetua a insercao
	query := `INSERT INTO log (
                          cod_log
                        , descricao_log
						, cod_recurso
						, criado_por_log
                 ) VALUES (?, ?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		msg = fmt.Sprintf("Erro ao preparar a query: %s", err.Error())
		return msg, err
	}

	res, err := stmt.Exec(novo_log.Codigo, novo_log.Descricao, novo_log.Codigo_recurso, novo_log.Criado_por)
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

func Log_Consultar_Codigo(codigo_recurso string) (models.Log, bool, error, string) {
	var msg string

	var log models.Log
	db, err := Conectar()
	if err != nil {
		msg = fmt.Sprintf("Erro ao conectar: %s", err.Error())
		return log, false, err, msg
	}
	defer db.Close()

	query := "SELECT cod_log, descricao_log, criado_por_log, data_ult_atu_servico FROM servico WHERE cod_recurso = ?"

	rows, err := db.Query(query, codigo_recurso)
	if err != nil {
		return log, false, err, err.Error()
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o código de recurso: %s ", codigo_recurso)
		return log, false, nil, msg
	}

	err = rows.Scan(&log.Codigo, &log.Descricao, &log.Criado_por, &log.Data_ult_atu)
	if err != nil {
		return log, false, err, err.Error() // Erro real
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Sucesso - Log do recurso %s encontrado com sucesso", codigo_recurso)
	return log, true, nil, msg
}
