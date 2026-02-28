package controllers

import (
	"api-go-crud/src/database"

	"github.com/gofiber/fiber/v2"
)

func Consulta_Log_Codigo(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	log, achou, err, msg := database.Log_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Log não existe...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Log existe...
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    log, // Adiciona o objeto inteiro aqui
		})
	}
}
