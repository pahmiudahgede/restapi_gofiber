package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"rijik.id/restapi_gofiber/dto"
	"rijik.id/restapi_gofiber/internal/middleware"
	"rijik.id/restapi_gofiber/internal/repository"
	"rijik.id/restapi_gofiber/internal/service"
	"rijik.id/restapi_gofiber/internal/utils"
)

func RegisterRoutes(app *fiber.App) {

	app.Post("/register", func(c *fiber.Ctx) error {
		var userDTO dto.UserRegisterDTO
		if err := c.BodyParser(&userDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(fiber.StatusBadRequest, "Invalid request"))
		}

		log.Printf("Received registration request: %+v", userDTO)

		if err := service.RegisterUser(userDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(utils.NewSuccessResponse("User registered successfully", nil))
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var userDTO dto.UserLoginDTO
		if err := c.BodyParser(&userDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(fiber.StatusBadRequest, "Invalid request"))
		}

		token, err := service.LoginUser(userDTO)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.NewErrorResponse(fiber.StatusUnauthorized, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(utils.NewSuccessResponse("Login successful", map[string]string{
			"access_token": token,
		}))
	})

	app.Get("/user", middleware.AuthMiddleware, func(c *fiber.Ctx) error {

		userID := c.Locals("user")

		user, err := repository.GetUserByID(userID.(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve user"))
		}

		return c.Status(fiber.StatusOK).JSON(utils.NewSuccessResponse("User data fetched successfully", user))
	})

}
