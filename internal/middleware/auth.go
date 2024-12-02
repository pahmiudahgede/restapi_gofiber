package middleware

import (
	"errors"
	"strings"

	"rijik.id/restapi_gofiber/internal/config"
	"rijik.id/restapi_gofiber/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewErrorResponse(fiber.StatusUnauthorized, "Missing Authorization header"))
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewErrorResponse(fiber.StatusUnauthorized, "Invalid Authorization format"))
	}

	tokenString := bearerToken[1]

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewErrorResponse(fiber.StatusUnauthorized, "Invalid or expired token"))
	}

	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.NewErrorResponse(fiber.StatusUnauthorized, "Invalid token claims"))
	}

	c.Locals("user", userID)
	return c.Next()
}
