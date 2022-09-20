package product

import (
	"github.com/gofiber/fiber/v2"
)

func Delete(h *ProductHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		id := ctx.Params("id")

		if err := h.ProductUsecase.Delete(id); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error could not create product " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).Send(nil)
	}
}
