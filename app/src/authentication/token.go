package authentication

import (
	"api-go-crud/src/configs"
	"time"

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
		return "", err // Erro ao gerar token
	}
	return token, nil
}
