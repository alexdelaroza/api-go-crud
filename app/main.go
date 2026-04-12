package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	config "api-go-crud/src/configs"
	"api-go-crud/src/router"
)

func main() {
	//database.Create_table()

	// Carregar as Variavei de Ambiente
	config.CarregarConfig()

	// cria a instancia do WEB server
	app := fiber.New()

	// CORS é uma medida de segurança para proteger os usuários contra vulnerabilidades e ataques maliciosos.
	// AllowCredentials definida como true, permite que o servidor inclua cookies e cabeçalhos de autenticação na solicitação.
	app.Use(cors.New(cors.Config{
		// Permite tanto o acesso local quanto o acesso pelo IP da sua rede
		AllowOrigins:     "http://localhost:3001, http://192.168.31.14:3001",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// setup app routes
	router.Setup(app)

	// iniciamos o seridor
	fmt.Println("Escutando na Porta:", config.Porta)
	app.Listen(fmt.Sprintf(":%d", config.Porta))

}
