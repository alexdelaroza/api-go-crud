package controllers

import (
	database "api-go-crud/src/databases"
	"api-go-crud/src/models"
	"api-go-crud/src/validation"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CRUD - Usuarios
func InserirUsuarios(c *fiber.Ctx) error {
	var novo_usuario models.Usuario_input

	err := c.BodyParser(&novo_usuario)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	// Como a função pede um ponteiro (*models.Usuario_input), você usa o símbolo "&"" na frente da variável.
	// & <-- passa o "endereço" da variável
	valido, msg_ret_ent := validation.ValidarInputUsuarios(novo_usuario)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg_ret_ent,
		})
	}

	//Valida se o Email ja existe no Cadastro
	achou_email, msg_ret_email, err := database.Usuario_Consultar_Email(novo_usuario.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": msg_ret_email,
			"error":   err.Error(),
		})
	}
	if achou_email {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg_ret_email + " - Não será inserido",
		})
	}

	// Valida se o Login ja existe no Cadastro
	achou, msg_ret_login, err := database.Usuario_Consultar_Login(novo_usuario.Login)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": msg_ret_login,
			"error":   err.Error(),
		})
	}
	if achou {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg_ret_login + " - Não será inserido",
		})
	}

	// Usuário não existe -> INSERIR
	retorno_id, msg, err := database.Usuario_Inserir(novo_usuario)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	// log -> INSERIR
	var novo_log models.Log_input
	novo_log.Codigo_recurso = strconv.Itoa(retorno_id)
	novo_log.Criado_por = "1"
	novo_log.Descricao = "insercao de usuario"

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

}

func AtualizarUsuarios(c *fiber.Ctx) error {
	var altera_usuario models.Usuario_input
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
	valido, msg_ret := validation.ValidarId(id)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	// Valida se o Email ja existe no Cadastro
	achou_email, msg_ret_email, err := database.Usuario_Consultar_Email(altera_usuario.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": msg_ret_email,
			"error":   err.Error(),
		})
	}
	if achou_email {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg_ret_email + " - Não será inserido",
		})
	}

	// Valida se o Login ja existe no Cadastro
	achou_login, msg_ret_login, err := database.Usuario_Consultar_Login(altera_usuario.Login)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": msg_ret_login,
			"error":   err.Error(),
		})
	}
	if achou_login {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg_ret_login + " - Não será inserido",
		})
	}

	// Valida se o Usuario existe no Cadastro
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
		msg, err := database.Usuario_Atualizar(id, altera_usuario)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		// log -> INSERIR
		var novo_log models.Log_input
		novo_log.Codigo_recurso = id
		novo_log.Criado_por = "1"
		novo_log.Descricao = "alteracao de usuario"

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
			"user":    altera_usuario,
		})
	}
}

func DeletarUsuarios(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")

	// Valida Dados de Entrada
	valido, msg_ret := validation.ValidarId(id)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
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

		// log -> INSERIR
		var novo_log models.Log_input
		novo_log.Codigo_recurso = id
		novo_log.Criado_por = "1"
		novo_log.Descricao = "delecao de usuario"

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

func ListarUsuarios(c *fiber.Ctx) error {
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

func ConsultarCodigoUsuarios(c *fiber.Ctx) error {
	var id string
	id = c.Params("id")
	userID := c.Locals("user_id")

	fmt.Println("ConsultarCodigoUsuarios = :", userID)

	// Valida Dados de Entrada
	valido, msg_ret := validation.ValidarId(id)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
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
