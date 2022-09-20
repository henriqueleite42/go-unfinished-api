package usecase

import (
	"errors"
	"unfinished-api/src/models/entities"
	"unfinished-api/src/models/files"
)

type productUsecase struct {
	productRepo entities.ProductRepository
	imageRepo   files.ImageRepository
}

func NewProductUsecase(s entities.ProductRepository, i files.ImageRepository) entities.ProductUsecase {
	return &productUsecase{
		productRepo: s,
		imageRepo:   i,
	}
}

func (uc *productUsecase) GetProductsList(p entities.ProductGetByStoreInput) ([]entities.Product, error) {
	return uc.productRepo.GetProductsList(p)
}

func (uc *productUsecase) Create(p *entities.Product) error {
	if p.Image != "" {
		imageUrl, err := uc.imageRepo.CreateFromURL(p.Image)
		if err != nil {
			return errors.New("fail to save image")
		}

		p.Image = imageUrl
	}

	return uc.productRepo.Create(p)
}

func (uc *productUsecase) Delete(ID string) error {
	return uc.productRepo.Delete(ID)
}
