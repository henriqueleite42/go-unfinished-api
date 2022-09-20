package client

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func Create(h *ClientHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		client := entities.Client{}
		err := ctx.BodyParser(&client)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error parsing json " + err.Error(),
			})
		}

		if err = h.ClientUsecase.Create(&client); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not create client " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(client)
	}
}
