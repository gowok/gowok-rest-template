package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gowok/gowok-rest-template/module/auth/dto"
	"github.com/gowok/ioc"
)

type Http struct {
	service *Service
}

func NewHttp() *Http {
	return &Http{
		ioc.Get(Service{}),
	}
}

func InitHttp() {
	app := ioc.Get(fiber.App{})
	auth := app.Group("/auth")

	h := NewHttp()
	auth.Post("/login", h.Login)
}

func (h Http) Login(c *fiber.Ctx) error {
	input := dto.LoginReq{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.service.Login(c.Context(), input.Email, input.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "mantap",
	})
}
