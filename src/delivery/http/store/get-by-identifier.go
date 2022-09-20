package store

import (
	"net/http"
	"strings"

	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

func GetByIdentifier(h *StoreHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		identifier := ctx.Params("identifier")

		var store entities.Store
		var err error
		if strings.Contains(identifier, "#") {
			store, err = h.StoreUsecase.GetByTag(identifier)
		} else {
			store, err = h.StoreUsecase.GetByID(identifier)
		}

		if err != nil {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(store)
	}
}
