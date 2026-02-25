package controllers

import (
	"api-go-crud/src/database"
	"api-go-crud/src/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CRUD - Usuarios
func Insere_Usuario(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// Valida Dados de Entrada
	if data["codigo"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	if data["nome"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'nome' é obrigatório e deve ser preenchido!",
		})
	}

	if data["login"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'login' é obrigatório e deve ser preenchido!",
		})
	}

	if data["senha"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'senha' é obrigatório e deve ser preenchido!",
		})
	}

	if data["email"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'email' é obrigatório e deve ser preenchido!",
		})
	}

	if data["tipo"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'tipo' é obrigatório e deve ser preenchido!",
		})
	}

	var novo_usuario models.Usuario
	novo_usuario.Codigo = data["codigo"]
	novo_usuario.Nome = data["nome"]
	novo_usuario.Login = data["login"]
	novo_usuario.Senha = data["senha"]
	novo_usuario.Email = data["email"]
	novo_usuario.Tipo = data["tipo"]

	// Verificar se o Usuario ja existe no Cadastro
	usuario, achou, err := database.Usuario_Consultar_Codigo(novo_usuario.Codigo)
	if err != nil {
		return c.Status(500).SendString("Erro interno no banco")
	}

	if !achou {
		// Usuário não existe. Seguindo para inserção...
		msg, err := database.Usuario_Inserir(novo_usuario)

		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(201)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    novo_usuario, // Adiciona o objeto inteiro aqui
		})
	} else {
		// Usuário existe. Não sera inserido...
		var msg string
		msg = fmt.Sprintf("Usuário %s encontrado...", usuario.Codigo)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    novo_usuario, // Adiciona o objeto inteiro aqui
		})
	}

	//return c.JSON(novo_usuario)
}

func Atualiza_Usuario(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	var altera_usuario models.Usuario
	altera_usuario.Codigo = id
	altera_usuario.Nome = data["nome"]
	altera_usuario.Login = data["login"]
	altera_usuario.Senha = data["senha"]
	altera_usuario.Email = data["email"]
	altera_usuario.Tipo = data["tipo"]

	// Verificar se o Usuario ja existe no Cadastro
	usuario, achou, err := database.Usuario_Consultar_Codigo(altera_usuario.Codigo)
	if err != nil {
		return c.Status(500).SendString("Erro interno no banco")
	}

	if !achou {
		// Usuário não existe. Não sera alterado...
		var msg string
		msg = fmt.Sprintf("Usuário %s não encontrado...", altera_usuario.Codigo)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    usuario, // Adiciona o objeto inteiro aqui
		})

	} else {
		// Usuário existe. Seguindo para alteração...
		msg, err := database.Usuario_Atualizar(altera_usuario)

		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(201)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    altera_usuario, // Adiciona o objeto inteiro aqui
		})
	}

	//return c.JSON(altera_usuario)
}

func Deleta_Usuario(c *fiber.Ctx) error {
	var codigo_usuario string
	codigo_usuario = c.Params("id")

	// Valida Dados de Entrada
	if codigo_usuario == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Usuario ja existe no Cadastro
	_, achou, err := database.Usuario_Consultar_Codigo(codigo_usuario)
	if err != nil {
		return c.Status(500).SendString("Erro interno no banco")
	}

	if !achou {
		// Usuário não existe. Não é possivel efetuar a exclusão...
		var msg string
		msg = fmt.Sprintf("Usuário %s não encontrado...", codigo_usuario)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    codigo_usuario,
		})

	} else {
		// Usuário existe. E sera efetuada a exclusão...
		msg, err := database.Usuario_Deletar(codigo_usuario)

		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(201)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    codigo_usuario, 
		})
	}

}

func Consulta_Usuario(c *fiber.Ctx) error {
	lista, err := database.Usuario_Consultar()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar usuários"})
	}

	c.Status(201)
	// Retorna a lista completa de Usuarios Cadastrados como um array JSON
	return c.JSON(lista)

}

func Consulta_Usuario_Codigo(c *fiber.Ctx) error {
	var data models.Usuario
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Usuario ja existe no Cadastro
	usuario, achou, err := database.Usuario_Consultar_Codigo(id)
	if err != nil {
		return c.Status(500).SendString("Erro interno no banco")
	}

	if !achou {
		// Usuário não existe...
		var msg string
		msg = fmt.Sprintf("Usuário %s não encontrado...", usuario.Codigo)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    usuario, // Adiciona o objeto inteiro aqui
		})
	} else {
		// Usuário existe...
		var msg string
		msg = fmt.Sprintf("Usuário %s encontrado...", usuario.Codigo)
		c.Status(201)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    usuario, // Adiciona o objeto inteiro aqui
		})
	}

	//return c.JSON(novo_usuario)
}
