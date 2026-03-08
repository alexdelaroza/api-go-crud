package controllers

import (
	"api-go-crud/src/database"
	"api-go-crud/src/models"

	"github.com/gofiber/fiber/v2"
)

func Valida_input(Descricao string, Valor float64) (bool, string) {
	if Descricao == "" {
		return false, "O campo 'descricao' é obrigatório e deve ser preenchido!"
	}

	if Valor == 0 {
		return false, "O campo 'valor' é obrigatório e deve ser preenchido!"
	}
	return true, ""
}

// CRUD - Servico
func Insere_Servico(c *fiber.Ctx) error {
	var novo_servico models.Servico_input

	err := c.BodyParser(&novo_servico)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	valido, msg_ret := Valida_input(novo_servico.Descricao, novo_servico.Valor)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	// Verificar se o Servico ja existe no Cadastro
	achou, msg_ret, err := database.Servico_Consultar_Descricao(novo_servico.Descricao)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": msg_ret,
			"error":   err.Error(),
		})
	}

	if !achou {
		// Servico não existe -> INSERIR
		msg, err := database.Servico_Inserir(novo_servico)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"msg":   msg,
				"error": err.Error(),
			})
		}

		// log -> INSERIR
		var novo_log models.Log_input
		// busca ultimo id serviço
		ok, retorno, err := database.Servico_ultimo_id()
		if !ok {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"retorno": retorno,
				"error":   err.Error(),
			})
		}
		novo_log.Codigo_recurso = retorno
		novo_log.Criado_por = "1"
		novo_log.Descricao = "insercao de servico"

		msg, err = database.Log_Inserir(novo_log)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": msg,
				"error":   err.Error(),
			})
		}
		// log -> INSERIR

		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	} else {
		// Servico existe. Não sera inserido...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg_ret + " - Não será inserido",
		})
	}
}

func Atualiza_Servico(c *fiber.Ctx) error {
	var altera_servico models.Servico_input

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
		msg, err := database.Servico_Atualizar(id, altera_servico)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": msg,
				"error":   err.Error(),
			})
		}

		// log -> INSERIR
		var novo_log models.Log_input
		novo_log.Codigo_recurso = id
		novo_log.Criado_por = "1"
		novo_log.Descricao = "alteracao de servico"

		msg, err = database.Log_Inserir(novo_log)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": msg,
				"error":   err.Error(),
			})
		}
		// log -> INSERIR

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
		// Servico não existe. Não é possivel efetuar a exclusão...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    id,
		})
	} else {
		// Servico existe -> exclusão...
		msg, err := database.Servico_Deletar(id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		// log -> INSERIR
		var novo_log models.Log_input
		novo_log.Codigo_recurso = id
		novo_log.Criado_por = "1"
		novo_log.Descricao = "delecao de servico"

		msg, err = database.Log_Inserir(novo_log)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": msg,
				"error":   err.Error(),
			})
		}
		// log -> INSERIR

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
		"message":  msg,
		"services": lista,
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
			"service": servico, // Adiciona o objeto inteiro aqui
		})
	}
}
