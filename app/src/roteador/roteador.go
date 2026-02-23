package roteador

import (
	"api-go-crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/usuario", controllers.Insere_Usuario)
	app.Put("/usuario/id", controllers.Atualiza_Usuario)
	app.Delete("/usuario/:id", controllers.Deleta_Usuario)
	app.Get("/usuario", controllers.Consulta_Usuario)
	app.Get("/usuario/:id", controllers.Consulta_Usuario)

	app.Post("/servico", controllers.Insere_Servico)
	app.Put("/servico/:id", controllers.Atualiza_Servico)
	app.Delete("/servico/:id", controllers.Deleta_Servico)
	app.Get("/servico", controllers.Consulta_Servico)
	app.Get("/servico/:id", controllers.Consulta_Servico)
}
