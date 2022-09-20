package usecase

import (
	"unfinished-api/src/models/entities"
)

type clientUsecase struct {
	clientRepo entities.ClientRepository
}

func NewClientUsecase(c entities.ClientRepository) entities.ClientUsecase {
	return &clientUsecase{
		clientRepo: c,
	}
}

func (cUc *clientUsecase) GetByID(discordID string) (entities.Client, error) {
	return cUc.clientRepo.GetByID(discordID)
}

func (cUc *clientUsecase) Update(c *entities.Client) error {
	return cUc.clientRepo.Update(c)
}

func (cUc *clientUsecase) Create(c *entities.Client) error {
	return cUc.clientRepo.Create(c)
}
