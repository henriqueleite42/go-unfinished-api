package product

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductUsecase entities.ProductUsecase
}

func NewHandler(router *fiber.App, sUc entities.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: sUc,
	}

	router.Get("/products", GetProductsList(handler))
	router.Post("/products", Create(handler))
	router.Delete("/products/:id", Delete(handler))
}
