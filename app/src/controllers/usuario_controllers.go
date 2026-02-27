package controllers

import (
	"api-go-crud/src/database"
	"api-go-crud/src/models"

	"github.com/gofiber/fiber/v2"
)

// CRUD - Usuarios
func Insere_Usuario(c *fiber.Ctx) error {
	var novo_usuario models.Usuario

	err := c.BodyParser(&novo_usuario)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	if novo_usuario.Codigo == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_usuario.Nome == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'nome' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_usuario.Login == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'login' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_usuario.Senha == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'senha' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_usuario.Email == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'email' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_usuario.Tipo == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'tipo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Usuario ja existe no Cadastro
	_, achou, err, msg := database.Usuario_Consultar_Codigo(novo_usuario.Codigo)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Usuário não existe -> INSERIR
		msg, err := database.Usuario_Inserir(novo_usuario)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Usuário existe. Não sera inserido...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}
}

func Atualiza_Usuario(c *fiber.Ctx) error {
	var altera_usuario models.Usuario

	err := c.BodyParser(&altera_usuario)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	var id string
	id = c.Params("id")
	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Usuario ja existe no Cadastro
	_, achou, err, msg := database.Usuario_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Usuário não existe. Não sera alterado...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})

	} else {
		// Usuário existe. Seguindo para alteração...
		altera_usuario.Codigo = id
		msg, err := database.Usuario_Atualizar(altera_usuario)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    altera_usuario,
		})
	}
}

func Deleta_Usuario(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Usuario ja existe no Cadastro
	_, achou, err, msg := database.Usuario_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Usuário não existe. Não é possivel efetuar a exclusão...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})
	} else {
		// Usuário existe. E sera efetuada a exclusão...
		msg, err := database.Usuario_Deletar(id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})
	}
}

func Consulta_Usuario(c *fiber.Ctx) error {
	lista, err, msg := database.Usuario_Consultar()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	// Retorna a lista completa de Usuarios Cadastrados como um array JSON
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": msg,
		"user":    lista,
	})
}

func Consulta_Usuario_Codigo(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	usuario, achou, err, msg := database.Usuario_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Usuário não existe...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Usuário existe...
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    usuario, // Adiciona o objeto inteiro aqui
		})
	}
}

// CRUD - Servico
func Insere_Servico(c *fiber.Ctx) error {
	var novo_servico models.Servico

	err := c.BodyParser(&novo_servico)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	if novo_servico.Codigo == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_servico.Descricao == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'descricao' é obrigatório e deve ser preenchido!",
		})
	}

	if novo_servico.Valor == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'valor' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Servico ja existe no Cadastro
	_, achou, err, msg := database.Servico_Consultar_Codigo(novo_servico.Codigo)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Servico não existe -> INSERIR
		msg, err := database.Servico_Inserir(novo_servico)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Servico existe. Não sera inserido...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}
}

func Atualiza_Servico(c *fiber.Ctx) error {
	var altera_servico models.Servico

	err := c.BodyParser(&altera_servico)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	var id string
	id = c.Params("id")
	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Servico ja existe no Cadastro
	_, achou, err, msg := database.Servico_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Servico não existe. Não sera alterado...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})
	} else {
		// Servico existe. Seguindo para alteração...
		altera_servico.Codigo = id
		msg, err := database.Servico_Atualizar(altera_servico)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    altera_servico,
		})
	}
}

func Deleta_Servico(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	// Verificar se o Servico ja existe no Cadastro
	_, achou, err, msg := database.Servico_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Usuário não existe. Não é possivel efetuar a exclusão...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})
	} else {
		// Usuário existe. E sera efetuada a exclusão...
		msg, err := database.Servico_Deletar(id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})
	}
}

func Consulta_Servico(c *fiber.Ctx) error {
	lista, err, msg := database.Servico_Consultar()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	// Retorna a lista completa de serviços Cadastrados como um array JSON
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": msg,
		"user":    lista,
	})
}

func Consulta_Servico_Codigo(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	if id == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "O campo 'codigo' é obrigatório e deve ser preenchido!",
		})
	}

	servico, achou, err, msg := database.Servico_Consultar_Codigo(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		// Servico não existe...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Servico existe...
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    servico, // Adiciona o objeto inteiro aqui
		})
	}
}
