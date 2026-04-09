package authentication

import (
	"api-go-crud/src/config"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CriarToken(usuarioID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": usuarioID,
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	}

	//config.JwtSecret => JwtSecret := []byte("minha_chave_secreta_123")

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(config.JwtSecret)
	if err != nil {
		return "Erro ao gerar token", err
	}
	return token, nil
}

func AuthRequired(c *fiber.Ctx) error {
	var JwtSecret = []byte("sua_chave_secreta_super_segura")
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
		return JwtSecret, nil
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
