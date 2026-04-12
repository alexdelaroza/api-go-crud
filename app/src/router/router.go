package router

import (
	"api-go-crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Login
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)

	// Ativa a autenticacao para as rotas abaixo
	//app.Use(controllers.AuthorizationCookie)
	app.Use(controllers.AuthorizationHeader)

	// Esta rota o React chamará para saber "quem sou eu" assim que abrir o site
	app.Get("/user", controllers.ObterUsuarioPeloToken)

	// Usuarios
	app.Post("/usuarios", controllers.InserirUsuarios)
	app.Put("/usuarios/:id", controllers.AtualizarUsuarios)
	app.Delete("/usuarios/:id", controllers.DeletarUsuarios)
	app.Get("/usuarios", controllers.ListarUsuarios)
	app.Get("/usuarios/:id", controllers.ConsultarCodigoUsuarios)

	// Servicos
	app.Post("/servicos", controllers.InserirServicos)
	app.Put("/servicos/:id", controllers.AtualizarServicos)
	app.Delete("/servicos/:id", controllers.DeletarServicos)
	app.Get("/servicos", controllers.ListarServicos)
	app.Get("/servicos/:id", controllers.ConsultarCodigoServicos)

	// Log
	app.Get("/logs", controllers.Consulta_Log)
}
