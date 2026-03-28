package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"api-go-crud/src/config"
	"api-go-crud/src/router"
)

func main() {
	//database.Create_table()

	// Carregar as Variavei de Ambiente
	config.CarregarConfig()

	// cria a instancia do WEB server
	app := fiber.New()
	// setup app routes
	router.Setup(app)
	// iniciamos o seridor
	fmt.Println("Escutando na Porta:", config.Porta)
	app.Listen(fmt.Sprintf(":%d", config.Porta))

}
