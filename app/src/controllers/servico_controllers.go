package controllers

import (
	"api-go-crud/src/database"
	"api-go-crud/src/models"
	"api-go-crud/src/validation"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CRUD - Servico
func InserirServicos(c *fiber.Ctx) error {
	var novo_servico models.Servico_input

	err := c.BodyParser(&novo_servico)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	valido, msg_ret := validation.ValidarInputServicos(novo_servico)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	// Verificar se o Servico ja existe no Cadastro
	achou, msg_ret, err := database.ServicosConsultarDescricao(novo_servico.Descricao)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": msg_ret,
			"error":   err.Error(),
		})
	}

	if !achou {
		// Servico não existe -> INSERIR
		retorno_id, msg, err := database.ServicosInserir(novo_servico)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"msg":   msg,
				"error": err.Error(),
			})
		}

		// log -> INSERIR
		var novo_log models.Log_input
		novo_log.Codigo_recurso = strconv.Itoa(retorno_id)
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

func AtualizarServicos(c *fiber.Ctx) error {
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
	valido, msg_ret := validation.ValidarId(id)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	// Verificar se o Servico ja existe no Cadastro
	_, achou, err, msg := database.ServicosConsultarCodigo(id)
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
		msg, err := database.ServicosAtualizar(id, altera_servico)
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

func DeletarServicos(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")
	// Valida Dados de Entrada
	valido, msg_ret := validation.ValidarId(id)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	// Verificar se o Servico ja existe no Cadastro
	_, achou, err, msg := database.ServicosConsultarCodigo(id)
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
		msg, err := database.ServicosDeletar(id)
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

func ConsultarServicos(c *fiber.Ctx) error {
	lista, err, msg := database.ServicosConsultar()
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

func ConsultarCodigoServicos(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")
	// Valida Dados de Entrada
	valido, msg_ret := validation.ValidarId(id)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	servico, achou, err, msg := database.ServicosConsultarCodigo(id)
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
