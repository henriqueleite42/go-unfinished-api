package store

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func Update(h *StoreHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		ID := ctx.Params("id")

		store := entities.Store{}
		err := ctx.BodyParser(&store)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error parsing json " + err.Error(),
			})
		}

		store.ID = ID

		if err = h.StoreUsecase.Update(&store); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not update store " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(store)
	}
}
