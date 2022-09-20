package client

import (
	"unfinished-api/src/models/entities"

	"github.com/gofiber/fiber/v2"
)

type ClientHandler struct {
	ClientUsecase entities.ClientUsecase
}

func NewHandler(router *fiber.App, cUc entities.ClientUsecase) {
	handler := &ClientHandler{
		ClientUsecase: cUc,
	}

	router.Get("/clients/:discordId", GetByID(handler))
	router.Patch("/clients", Update(handler))
	router.Post("/clients", Create(handler))
}
