package database

import (
	"fmt"

	"api-go-crud/src/models"
)

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
