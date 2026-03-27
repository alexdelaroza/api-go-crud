package roteador

import (
	"api-go-crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Usuarios
	app.Post("/usuarios", controllers.InserirUsuarios)
	app.Put("/usuarios/:id", controllers.AtualizarUsuarios)
	app.Delete("/usuarios/:id", controllers.DeletarUsuarios)
	app.Get("/usuarios", controllers.ListarUsuarios)
	app.Get("/usuarios/:id", controllers.ConsultarCodigoUsuarios)

	// Login
	app.Post("/login", controllers.EfetuarLoginUsuarios)
	
	// Servicos
	app.Post("/servicos", controllers.InserirServicos)
	app.Put("/servicos/:id", controllers.AtualizarServicos)
	app.Delete("/servicos/:id", controllers.DeletarServicos)
	app.Get("/servicos", controllers.ConsultarServicos)
	app.Get("/servicos/:id", controllers.ConsultarCodigoServicos)

	// Log
	app.Get("/logs", controllers.Consulta_Log)
}
