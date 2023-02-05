package rest

import (
	"github.com/gofiber/fiber/v2"
)

func Index(api *APIRest) {
	router := api.HTTP.Group("/")
	router.Get("", api.Index)
}

func (api APIRest) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "welcome to gowok",
	})
}
