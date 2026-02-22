package controllers

import (
	"api-go-crud/src/database"
	"api-go-crud/src/models"

	"github.com/gofiber/fiber/v2"
)

// TESTE
func Consulta_Usuario(c *fiber.Ctx) error {
	return c.SendString("Hello Consulta_Usuario, World 👋!")
}

// TESTE
func Insere_Usuario(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	if data["tipo"] != data["confirm_tipo"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "tipo do not match!",
		})
	}

	var novo_usuario models.Usuario
	novo_usuario.Codigo = data["codigo"]
	novo_usuario.Nome = data["nome"]
	novo_usuario.Login = data["login"]
	novo_usuario.Senha = data["senha"]
	novo_usuario.Email = data["email"]
	novo_usuario.Tipo = data["tipo"]

	database.Usuario_Inserir(novo_usuario)

	return c.JSON(novo_usuario)
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
