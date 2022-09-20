package store

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
	StoreUsecase entities.StoreUsecase
}

func NewHandler(router *fiber.App, sUc entities.StoreUsecase) {
	handler := &StoreHandler{
		StoreUsecase: sUc,
	}

	router.Get("/stores/:identifier", GetByIdentifier(handler))
	router.Patch("/stores", Update(handler))
	router.Post("/stores", Create(handler))
}
