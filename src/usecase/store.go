package usecase

import (
	"errors"
	"unfinished-api/src/models/entities"
	"unfinished-api/src/models/files"
)

type storeUsecase struct {
	storeRepo entities.StoreRepository
	imageRepo files.ImageRepository
}

func NewStoreUsecase(s entities.StoreRepository, i files.ImageRepository) entities.StoreUsecase {
	return &storeUsecase{
		storeRepo: s,
		imageRepo: i,
	}
}

func (uc *storeUsecase) GetByID(discordID string) (entities.Store, error) {
	return uc.storeRepo.GetByID(discordID)
}

func (uc *storeUsecase) GetByTag(discordTag string) (entities.Store, error) {
	return uc.storeRepo.GetByTag(discordTag)
}

func (uc *storeUsecase) Update(s *entities.Store) error {
	if s.Image != "" {
		imageUrl, err := uc.imageRepo.CreateFromURL(s.Image)
		if err != nil {
			return errors.New("fail to save image")
		}

		s.Image = imageUrl
	}

	return uc.storeRepo.Update(s)
}

func (uc *storeUsecase) Create(s *entities.Store) error {
	if s.Image != "" {
		imageUrl, err := uc.imageRepo.CreateFromURL(s.Image)
		if err != nil {
			return errors.New("fail to save image")
		}

		s.Image = imageUrl
	}

	return uc.storeRepo.Create(s)
}
