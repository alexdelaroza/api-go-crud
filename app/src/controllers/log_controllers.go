package controllers

import (
	"api-go-crud/src/database"

	"github.com/gofiber/fiber/v2"
)

func Consulta_Log(c *fiber.Ctx) error {
	id := c.Query("id")                 // ex: ?id=1
	dataInicio := c.Query("dataInicio") // ex: ?data=2026-03-08
	dataFim := c.Query("dataFim")       // ex: ?data=2026-03-08

	// Verifica entrada de dados
	if dataInicio == "" && dataFim == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "Informe 'dataInicio' e 'dataFim' na pesquisa - obrigatório e devem ser preenchido!",
		})
	}

	lista, achou, err, msg := database.Log_Consultar(id, dataInicio, dataFim)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Log não existe...
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Log existe...
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"log":     lista,
		})
	}

}
