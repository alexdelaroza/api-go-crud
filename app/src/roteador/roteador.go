package roteador

import (
	"api-go-crud/src/api"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/usuario", api.Consulta_Usuario)
	app.Post("/usuario", api.Insere_Usuario)

	app.Get("/usuario/:id", api.Consulta_Usuario)
	app.Put("/usuario/:id", api.Atualiza_Usuario)
	app.Delete("/usuario/:id", api.Deleta_Usuario)

	app.Get("/servico", api.Consulta_Servico)
	app.Post("/servico", api.Insere_Servico)

	app.Get("/servico/:id", api.Consulta_Servico)
	app.Put("/servico/:id", api.Atualiza_Servico)
	app.Delete("/servico/:id", api.Deleta_Servico)
}
