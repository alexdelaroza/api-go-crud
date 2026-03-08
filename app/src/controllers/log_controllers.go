package controllers

import (
	"api-go-crud/src/database"

	"github.com/gofiber/fiber/v2"
)

func Consulta_Log(c *fiber.Ctx) error {
	id := c.Query("id")     // ex: ?id=1
	data := c.Query("data") // ex: ?data=2026-03-08

	//busca log por ID
	if id != "" {
		lista, achou, err, msg := database.Log_Consultar_Codigo(id)
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

	//busca lista logs por data
	if data != "" {
		lista, achou, err, msg := database.Log_Consultar_Data(data)
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
			// Retorna a lista de logs filtrados por data como um array JSON
			c.Status(fiber.StatusOK)
			return c.JSON(fiber.Map{
				"message": msg,
				"logs":    lista,
			})
		}
	}

	// Se não enviou nada: retorna erro ou lista tudo
	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Map{
		"error": "Informe 'id' ou 'data' na pesquisa",
	})
}
