package api

import (
	"github.com/gofiber/fiber/v2"
)

// TESTE
func Consulta_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello Consulta_Usuario, World 👋!")
}

// TESTE
func Insere_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello Insere_Usuario, World 👋!")
}

// TESTE
func Atualiza_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello Atualiza_Usuario, World 👋!")
}

// TESTE
func Deleta_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello Deleta_Usuario, World 👋!")
}

// TESTE
func Consulta_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello Consulta_Servico, World 👋!")
}

// TESTE
func Insere_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello Insere_Servico, World 👋!")
}

// TESTE
func Atualiza_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello Atualiza_Servico, World 👋!")
}

// TESTE
func Deleta_Servico(c *fiber.Ctx) error {
	return c.SendString("Hello Deleta_Servico, World 👋!")
}

// TESTE
func Consulta_Log(c *fiber.Ctx) error {
	return c.SendString("Hello Consulta_Log, World 👋!")
}

// TESTE
func Insere_Log(c *fiber.Ctx) error {
	return c.SendString("Hello Insere_Log, World 👋!")
}
