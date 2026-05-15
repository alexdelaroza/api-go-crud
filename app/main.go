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

	config.CarregarConfig()

	err := database.ConectarDb()
	if err != nil {
		log.Fatal("Erro fatal ao conectar no banco:", err)
	}
	defer database.DB.Close()

	fmt.Println("Conexão com o banco estabelecida com sucesso!")

	//cria a instancia do WEB server
	app := fiber.New()

	app.Use(cors.New(cors.Config{
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
