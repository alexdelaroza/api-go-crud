package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	roteador "github.com/alexdelaroza/api-go-crud/src/roteador"
	usuario_Sql "github.com/alexdelaroza/api-go-crud/src/sql/usuario_Sql"
)

func main() {
	fmt.Println("alex")
	//usuario_Sql.Deleta()
	//usuario_Sql.Insere()
	//usuario_Sql.Atualiza()
	usuario_Sql.Consulta()

	// cria a instancia do WEB server
	app := fiber.New()
	// setup app routes
	roteador.Setup(app)
	// iniciamos o seridor
	app.Listen(":3000")
}
