package client

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetByID(h *ClientHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		discordID := ctx.Params("discordId")

		client, err := h.ClientUsecase.GetByID(discordID)
		if err != nil {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(client)
	}
}
