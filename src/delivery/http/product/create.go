package product

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func Create(h *ProductHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		product := entities.Product{}
		err := ctx.BodyParser(&product)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error parsing json " + err.Error(),
			})
		}

		if err = h.ProductUsecase.Create(&product); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not create product " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(product)
	}
}
