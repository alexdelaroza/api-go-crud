package controllers

import (
	"api-go-crud/src/authentication"
	config "api-go-crud/src/configs"
	database "api-go-crud/src/databases"
	"api-go-crud/src/models"
	"api-go-crud/src/validation"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizationHeader(c *fiber.Ctx) error {
	//config.JwtSecret = > var JwtSecret = []byte("sua_chave_secreta_super_segura")

	// 1. Obtém o cabeçalho Authorization
	authHeader := c.Get("Authorization")

	// 2. Verifica se o header está presente e no formato "Bearer <token>"
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token não fornecido ou formato inválido (Use: Bearer <token>)",
		})
	}

	// 3. Extrai apenas a string do token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// 4. Faz o Parse e a Validação do Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validação importante: garante que o método de assinatura é HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return config.JwtSecret, nil
	})

	// 5. Verifica se houve erro ou se o token é inválido/expirado
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Acesso negado: token inválido ou expirado",
		})
	}

	// 6. (Opcional) Extrai os dados do token (claims) e salva no contexto
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("user_id", claims["user_id"])
	}

	// 7. Se tudo estiver certo, segue para a próxima função (Handler)
	return c.Next()
}

func AuthorizationCookie(c *fiber.Ctx) error {
	type CustomClaims struct {
		UserID string `json:"user_id"`
		jwt.RegisteredClaims
	}

	//config.JwtSecret = > var JwtSecret = []byte("sua_chave_secreta_super_segura")

	// 1. Obtém o cookie
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated (cookie não encontrado)",
		})
	}

	// 2. Faz o Parse e a Validação do Token
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validação importante: garante que o método de assinatura é HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return config.JwtSecret, nil
	})

	// 3. Validação do token
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated (token inválido)",
		})
	}

	// 4. Extração dos dados
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		// Note que claims é um ponteiro para um mapa, acessamos assim:
		c.Locals("user_id", (*claims)["user_id"])
	}

	return c.Next()
}

func Logout(c *fiber.Ctx) error {
	// Cria um cookie expirado
	cookie := authentication.CriarCookieExpirado()

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout realizado com sucesso",
	})
}

func Login(c *fiber.Ctx) error {
	var login_usuario models.Usuario_login

	err := c.BodyParser(&login_usuario)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Valida Dados de Entrada
	valido, msg_ret := validation.ValidaLoginUsuarios(login_usuario)
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
	Token_retorno, err := authentication.CriarToken(usuarioID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": Token_retorno,
			"error":   err.Error(),
		})
	}

	// Criando Cookie... => Foi armazenado o token em um cookie com data de expiração.
	cookie := authentication.CriarCookieValido(Token_retorno)
	c.Cookie(&cookie)
	// Criando Cookie...

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": msg,
		"token":   Token_retorno, // Adiciona o objeto inteiro aqui
	})

}

func ObterUsuarioPeloToken(c *fiber.Ctx) error {
	// 1. Pega o ID que o seu Middleware AuthorizationHeader salvou no Locals
	// Fazemos a conversão para string (type assertion)
	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "ID do usuário não encontrado no token",
		})
	}

	// 2. Chama diretamente sua função de banco que já existe
	usuario, achou, err, msg := database.Usuario_Consultar_Codigo(userID)
	
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if !achou {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Usuário do token não encontrado no banco",
		})
	}

	// 3. Retorna o sucesso
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msg,
		"user":    usuario,
	})
}
