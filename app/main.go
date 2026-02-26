package main

import (
	"github.com/gofiber/fiber/v2"

	"api-go-crud/src/roteador"
)

func main() {
	//database.Create_table()
	//database.Usuario_Deletar()
	//database.Usuario_Inserir()
	//database.Usuario_Atualizar()
	//database.Usuario_Consultar()

	// cria a instancia do WEB server
	app := fiber.New()
	// setup app routes
	roteador.Setup(app)
	// iniciamos o seridor
	app.Listen(":3000")
}
