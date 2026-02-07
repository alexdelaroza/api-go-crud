package main

import (
	"fmt"

	api "github.com/alexdelaroza/api-go-crud/src/modulos/api"
	usuario_Sql "github.com/alexdelaroza/api-go-crud/src/modulos/sql/usuario_Sql"
)

func main() {
	fmt.Println("alex")

	api.Teste()
	usuario_Sql.Insere()

}
