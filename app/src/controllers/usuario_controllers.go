package controllers

import (
	"api-go-crud/src/authentication"
	"api-go-crud/src/database"
	"api-go-crud/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Login - Usuarios
func Valida_usuario_login(usuario models.Usuario_login) (bool, string) {

	if usuario.Login == "" && usuario.Email == "" {
		return false, "O campo 'email' ou 'login' é obrigatório e deve ser preenchido!"
	}

	if usuario.Senha == "" {
		return false, "O campo 'senha' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

// Validacoes - Usuarios
func Valida_usuario_input(usuario models.Usuario_input) (bool, string) {
	if usuario.Nome == "" {
		return false, "O campo 'nome' é obrigatório e deve ser preenchido!"
	}

	if usuario.Login == "" {
		return false, "O campo 'login' é obrigatório e deve ser preenchido!"
	}

	if usuario.Senha == "" {
		return false, "O campo 'senha' é obrigatório e deve ser preenchido!"
	}

	if usuario.Email == "" {
		return false, "O campo 'email' é obrigatório e deve ser preenchido!"
	}

	if usuario.Tipo == "" {
		return false, "O campo 'tipo' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

func Valida_usuario_id(id string) (bool, string) {
	if id == "" {
		return false, "O campo 'id' é obrigatório e deve ser preenchido!"
	}

	return true, ""
}

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
	valido, msg_ret_ent := Valida_usuario_input(novo_usuario)
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
	valido, msg_ret := Valida_usuario_id(id)
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
	valido, msg_ret := Valida_usuario_id(id)
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

	// Valida Dados de Entrada
	valido, msg_ret := Valida_usuario_id(id)
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

func EfetuarLoginUsuarios(c *fiber.Ctx) error {
	var login_usuario models.Usuario_login

	err := c.BodyParser(&login_usuario)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	valido, msg_ret := Valida_usuario_login(login_usuario)
	if !valido {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": msg_ret})
	}

	// Valida Usuario e Senha no banco de dados
	achou, msg, usuarioID, err := database.Usuario_Efetuar_Login(login_usuario)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !achou {
		// Login não é valido...
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}

	// Login é valido...
	retorno, err := authentication.CriarToken(usuarioID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": retorno,
			"error":   err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": msg,
		"token":   retorno, // Adiciona o objeto inteiro aqui
	})

}
