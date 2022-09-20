package product

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func GetProductsList(h *ProductHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		params := entities.ProductGetByStoreInput{}
		err := ctx.QueryParser(&params)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error parsing json " + err.Error(),
			})
		}

		products, err := h.ProductUsecase.GetProductsList(params)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not create product " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(products)
	}
}
