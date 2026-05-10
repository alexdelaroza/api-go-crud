package authentication

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CriarCookieValido(token string) fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 5),
		HTTPOnly: true,
	}
	return cookie
}

func CriarCookieExpirado() fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Data no passado (expira agora)
		HTTPOnly: true,                       
		SameSite: "Lax",
	}
	return cookie
}
