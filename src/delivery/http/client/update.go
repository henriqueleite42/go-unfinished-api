package client

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func Update(h *ClientHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		ID := ctx.Params("id")

		client := entities.Client{}
		err := ctx.BodyParser(&client)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error parsing json " + err.Error(),
			})
		}

		client.ID = ID

		if err = h.ClientUsecase.Update(&client); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not update client " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(client)
	}
}
