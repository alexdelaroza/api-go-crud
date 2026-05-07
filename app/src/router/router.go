package router

import (
	"api-go-crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Rotas Públicas
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)

	// Middleware de Autenticação (Verifica se o token existe e é válido)
	app.Use(controllers.AuthorizationHeader)

	// Identificação do Usuário Logado
	app.Get("/user", controllers.ObterUsuarioPeloToken)

	// --- USUÁRIOS ---
	app.Get("/usuarios", controllers.CheckAuth("admin", "vendedor"), controllers.ListarUsuarios)
	app.Get("/usuarios/:id", controllers.CheckAuth("admin", "vendedor"), controllers.ConsultarCodigoUsuarios)
	app.Post("/usuarios", controllers.CheckAuth("admin"), controllers.InserirUsuarios)
	app.Put("/usuarios/:id", controllers.CheckAuth("admin"), controllers.AtualizarUsuarios)
	app.Delete("/usuarios/:id", controllers.CheckAuth("admin"), controllers.DeletarUsuarios)

	// --- SERVIÇOS ---
	app.Get("/servicos", controllers.CheckAuth("admin", "vendedor"), controllers.ListarServicos)
	app.Get("/servicos/:id", controllers.CheckAuth("admin", "vendedor"), controllers.ConsultarCodigoServicos)
	app.Post("/servicos", controllers.CheckAuth("admin", "vendedor"), controllers.InserirServicos)
	app.Put("/servicos/:id", controllers.CheckAuth("admin", "vendedor"), controllers.AtualizarServicos)
	app.Delete("/servicos/:id", controllers.CheckAuth("admin", "vendedor"), controllers.DeletarServicos)

	// --- LOGS ---
	app.Get("/logs", controllers.CheckAuth("admin"), controllers.Consulta_Log)
}
