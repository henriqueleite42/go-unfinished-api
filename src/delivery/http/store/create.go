package store

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func Create(h *StoreHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		store := entities.Store{}
		err := ctx.BodyParser(&store)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error parsing json " + err.Error(),
			})
		}

		if err = h.StoreUsecase.Create(&store); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not create store " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(store)
	}
}
