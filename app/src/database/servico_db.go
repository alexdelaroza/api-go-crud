package database

import (
	"fmt"

	"api-go-crud/src/models"
)

// Servicos
func ServicosInserir(novo_servico models.Servico_input) (int, string, error) {
	var msg string

	query := `INSERT INTO servico (
                          descricao
						, valor
                 ) VALUES (?, ?)`

	stmt, err := DB.Prepare(query)
	if err != nil {
		msg = fmt.Sprintf("Erro ao preparar a query: %s", err.Error())
		return 0, msg, err
	}

	res, err := stmt.Exec(novo_servico.Descricao, novo_servico.Valor)
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
	return int(id), msg, nil
}

func ServicosAtualizar(codigo string, altera_servico models.Servico_input) (string, error) {
	var msg string

	query := `update  servico 
	            set   descricao = ? 
				  ,   valor     = ?
                where codigo    = ?`

	stmt, err := DB.Prepare(query)
	if err != nil {
		msg = fmt.Sprintf("Erro ao preparar a query: %s", err.Error())
		return msg, err
	}
	
	res, err := stmt.Exec(altera_servico.Descricao, altera_servico.Valor, codigo)
	if err != nil {
		msg = fmt.Sprintf("Erro ao executar a insercao: %s", err.Error())
		return msg, err
	}

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada para ser retornada
	msg = fmt.Sprintf("Sucesso! %d linha(s) afetada(s).", linhas)
	return msg, nil
}

func ServicosDeletar(codigo_servico string) (string, error) {
	var msg string

	query := `delete from servico where codigo = ?`

	stmt, _ := DB.Prepare(query)

	res, _ := stmt.Exec(codigo_servico)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()

	// fmt.Sprintf cria a string formatada
	msg = fmt.Sprintf("Sucesso! %d linha(s) deletada(s).", linhas)
	return msg, nil
}

func ServicosConsultar() ([]models.Servico_output, error, string) {
	var msg string

	query := `SELECT codigo, descricao, valor, data_criacao_atu FROM servico`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()

	var servicos []models.Servico_output

	for rows.Next() {
		var s models.Servico_output
		err := rows.Scan(&s.Codigo, &s.Descricao, &s.Valor, &s.Data_criacao_atu)
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

func ServicosConsultarCodigo(codigo_servico string) (models.Servico_output, bool, error, string) {
	var msg string

	var servico models.Servico_output

	query := "SELECT codigo, descricao, valor, data_criacao_atu FROM servico WHERE codigo = ?"

	rows, err := DB.Query(query, codigo_servico)
	if err != nil {
		return servico, false, err, err.Error()
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o código: %s ", codigo_servico)
		return servico, false, nil, msg
	}

	err = rows.Scan(&servico.Codigo, &servico.Descricao, &servico.Valor, &servico.Data_criacao_atu)

	if err != nil {
		return servico, false, err, err.Error() // Erro real
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Sucesso - Servico %s encontrado", servico.Codigo)
	return servico, true, nil, msg
}

func ServicosConsultarDescricao(descricao_servico string) (bool, string, error) {
	var msg, codigo string

	query := "SELECT codigo FROM servico WHERE descricao = ?"

	rows, err := DB.Query(query, descricao_servico)
	if err != nil {
		return false, err.Error(), err
	}
	defer rows.Close()

	if !rows.Next() {
		msg = fmt.Sprintf("Nenhum registro encontrado para o servico: %s ", descricao_servico)
		return false, msg, nil
	}

	err = rows.Scan(&codigo)
	if err != nil {
		return false, err.Error(), err // Erro real
	}

	// Sucesso - Encontrou
	msg = fmt.Sprintf("Servico %s encontrado", codigo)
	return true, msg, nil
}
