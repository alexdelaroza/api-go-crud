package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"api-go-crud/src/config"
	"api-go-crud/src/database"
	"api-go-crud/src/router"
)

func main() {

	// 1. Carregar as Variavei de Ambiente
	config.CarregarConfig()

	// 2. Conecta ao banco de dados
	err := database.ConectarDb()
	if err != nil {
		log.Fatal("Erro fatal ao conectar no banco:", err)
	}
	defer database.DB.Close()

	fmt.Println("Conexão com o banco estabelecida com sucesso!")

	//Caso precise criar a estrutura do Banco de Dados
	//database.Create_table()
	
	// 3. cria a instancia do WEB server
	app := fiber.New()

	// CORS é uma medida de segurança para proteger os usuários contra vulnerabilidades e ataques maliciosos.
	// AllowCredentials definida como true, permite que o servidor inclua cookies e cabeçalhos de autenticação na solicitação.
	app.Use(cors.New(cors.Config{
		// Permite tanto o acesso local quanto o acesso pelo IP da sua rede
		AllowOrigins:     "http://localhost:3001, http://127.0.0.1:3001, http://192.168.31.14:3001",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// setup app routes
	router.Setup(app)

	// iniciamos o servidor
	fmt.Println("API escutando na Porta:", config.Porta)
	app.Listen(fmt.Sprintf(":%d", config.Porta))

}
