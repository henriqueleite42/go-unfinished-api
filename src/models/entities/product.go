package entities

import (
	"time"
	"unfinished-api/src/models/enums"
)

//
//
// ENTITIES
//
//

type Product struct {
	ID             string               `json:"id" gorm:"column:id;primaryKey;"`
	StoreID        string               `json:"storeId" gorm:"column:storeId;not null;"`
	Name           string               `json:"name" gorm:"column:name;not null;"`
	Description    string               `json:"description" gorm:"column:description;not null;"`
	Color          string               `json:"color" gorm:"column:color;"`
	Image          string               `json:"image" gorm:"column:image;"`
	Price          float32              `json:"price" gorm:"column:price;not null;"`
	Type           enums.ProductType    `json:"type" sql:"type:ENUM('PACK', 'AUDIO', 'VIDEO', 'PHOTO', 'CUSTOM_AUDIO', 'CUSTOM_VIDEO', 'CUSTOM_PHOTO', 'RENT_A_GIRLFRIEND', 'SEXY_VOICE_CALL', 'SEXY_VIDEO_CALL')" gorm:"column:type;not null;"`
	DeliveryMethod enums.DeliveryMethod `json:"deliveryMethod" sql:"type:ENUM('DISCORD', 'GOOGLE_DRIVE')" gorm:"column:deliveryMethod;not null;"`
	CreatedAt      time.Time            `json:"createdAt" gorm:"column:created_at;not null;"`

	Store Store `json:"store"`
}

//
//
// REPOSITORY
//
//

type ProductListInput struct {
	StoreID string `validate:"required"`
	Type    enums.ProductType
	Take    uint
	Skip    uint
}

type ProductCreateInput struct {
	StoreID        string               `json:"storeId" validate:"required"`
	Name           string               `json:"name" validate:"required"`
	Description    string               `json:"description" validate:"required"`
	Color          string               `json:"color"`
	Image          string               `json:"image"`
	Price          float32              `json:"price" validate:"required"`
	Type           enums.ProductType    `json:"type" validate:"required"`
	DeliveryMethod enums.DeliveryMethod `json:"deliveryMethod" validate:"required"`
}

type ProductRepository interface {
	List(out *[]Product, in *ProductListInput) error
	Create(out *Product, in *ProductCreateInput) error
	Delete(ID string) error
}

//
//
// USECASE
//
//

type ProductGetByStoreInput struct {
	StoreID string            `json:"id" validate:"required"`
	Type    enums.ProductType `json:"type" validate:"required"`
	Page    uint              `json:"page" validate:"required"`
}

type ProductGetBestSellersInput struct {
	StoreID string       `json:"id" validate:"required"`
	Period  enums.Period `json:"period"`
}

type ProductUsecase interface {
	GetByStore(out *[]Product, in *ProductGetByStoreInput) error
	GetBestSellers(out *[]Product, in *ProductGetBestSellersInput) error
	Create(out *Product, in *ProductCreateInput) error
	Delete(ID string) error
}
