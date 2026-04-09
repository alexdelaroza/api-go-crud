package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"api-go-crud/src/config"
	"api-go-crud/src/router"
)

func main() {
	//database.Create_table()

	// Carregar as Variavei de Ambiente
	config.CarregarConfig()

	// cria a instancia do WEB server
	app := fiber.New()

	// CORS é uma medida de segurança que ajuda a proteger os usuários de sites da web contra vulnerabilidades e ataques maliciosos.
	// AllowCredentials definida como true, permite que o servidor inclua cookies e cabeçalhos de autenticação na solicitação.
	//Se não tiver como true, o frontend não vai conseguir pegar o cookie.
	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://localhost:%d", config.Porta), // Ou a porta do seu front
		AllowCredentials: true,
	}))

	// setup app routes
	router.Setup(app)
	// iniciamos o seridor
	fmt.Println("Escutando na Porta:", config.Porta)
	app.Listen(fmt.Sprintf(":%d", config.Porta))

}
