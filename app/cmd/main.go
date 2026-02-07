package main

import (
	"fmt"

	api "github.com/alexdelaroza/api-go-crud/src/modulos/api"
	roteador "github.com/alexdelaroza/api-go-crud/src/modulos/roteador"
	usuario_Sql "github.com/alexdelaroza/api-go-crud/src/modulos/sql/usuario_Sql"
)

func main() {
	fmt.Println("alex")

	api.Teste()
	roteador.Roteador()
	//usuario_Sql.Deleta()
	//usuario_Sql.Insere()
	//usuario_Sql.Atualiza()
	usuario_Sql.Consulta()

}
