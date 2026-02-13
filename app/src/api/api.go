package api

import (
	"github.com/gofiber/fiber/v2"
)

// TESTE
func Consulta_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello C, World 👋!")
}
// TESTE
func Insere_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}
// TESTE
func Atualiza_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}
// TESTE
func Deleta_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}


// TESTE
func Consulta_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello C, World 👋!")
}
// TESTE
func Insere_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}
// TESTE
func Atualiza_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}
// TESTE
func Deleta_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}


// TESTE
func Consulta_Log(c *fiber.Ctx) error {
	return c.SendString("Hello C, World 👋!")
}
// TESTE
func Insere_Log(c *fiber.Ctx) error {
	return c.SendString("Hello R, World 👋!")
}