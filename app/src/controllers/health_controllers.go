package controllers

import (
	"api-go-crud/src/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

var StartTime time.Time

func init() {
	StartTime = time.Now()
}

func HealthCheck(c *fiber.Ctx) error {
	if database.DB == nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Conexão com o banco não inicializada",
		})
	}

	if err := database.DB.Ping(); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  "unhealthy",
			"message": "Banco de dados inacessível",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "ok",
		"message": "API e Banco de Dados operacionais",
		"uptime": time.Since(StartTime).Round(time.Second).String(),
	})
}
