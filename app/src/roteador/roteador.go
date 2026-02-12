package roteador

import (
	"github.com/alexdelaroza/api-go-crud/src/api"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/usuario", api.Consulta_Usuario)
	app.Post("/api/usuario", api.Insere_Usuario)

	app.Get("/api/usuario/id", api.Consulta_Usuario)
	app.Put("/api/usuario/id", api.Atualiza_Usuario)
	app.Delete("/api/usuario/id", api.Deleta_Usuario)

	app.Get("/api/servico", api.Consulta_Servico)
	app.Post("/api/servico", api.Insere_Servico)

}
