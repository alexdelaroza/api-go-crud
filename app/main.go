package main

import (
	"github.com/gofiber/fiber/v2"

	"api-go-crud/src/database"
	"api-go-crud/src/roteador"
)

func main() {
	//database.Create_table()
	_, retorno, _ := database.Servico_ultimo_id()
	println(retorno)

	// cria a instancia do WEB server
	app := fiber.New()
	// setup app routes
	roteador.Setup(app)
	// iniciamos o seridor
	app.Listen(":3000")
}
