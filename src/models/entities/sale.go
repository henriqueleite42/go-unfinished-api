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

type SaleDiscount struct {
	SaleID       string             `json:"saleId" gorm:"column:sale_id;primaryKey;"`
	CouponID     string             `json:"couponId" gorm:"column:coupon_id;not null;"`
	DiscountType enums.DiscountType `json:"discountType" gorm:"column:discount_type;not null;"`
	Amount       float32            `json:"amount" gorm:"column:amount;not null;"`

	Coupon Coupon `json:"coupon"`
}

type SaleProduct struct {
	SaleID    string  `json:"saleId" gorm:"column:sale_id;primaryKey;"`
	ProductID string  `json:"productId" gorm:"column:product_id;not null;"`
	Name      string  `json:"name" gorm:"column:name;not null;"`
	Type      string  `json:"type" gorm:"column:type;not null;"`
	Color     string  `json:"color" gorm:"column:color;not null;"`
	Price     float32 `json:"price" gorm:"column:price;not null;"`

	Product Product `json:"product"`
}

type Sale struct {
	ID       string `json:"id" gorm:"column:id;primaryKey;"`
	StoreID  string `json:"storeId" gorm:"column:store_id;not null;"`
	ClientID string `json:"clientId" gorm:"column:client_id;not null;"`

	FinalPrice float32          `json:"finalPrice" gorm:"column:final_price;not null;"`
	Origin     string           `json:"origin" gorm:"column:origin;not null;"`
	Status     enums.SaleStatus `json:"status" gorm:"column:status;not null;"`

	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;not null;"`
	FinishedAt time.Time `json:"finishedAt" gorm:"column:finished_at;not null;"`

	Store        Store        `json:"store"`
	Client       Client       `json:"client"`
	SaleProduct  SaleProduct  `gorm:"ForeignKey:SaleID"`
	SaleDiscount SaleDiscount `gorm:"ForeignKey:SaleID"`
}

//
//
// REPOSITORY
//
//

type SaleGetExpensesInput struct {
	ClientID string       `json:"clientId" validate:"required"`
	Period   enums.Period `json:"period"`
}

type SaleGetExpensesOutput struct {
	ByType         map[enums.ProductType]float32 `json:"byType"`
	ByStore        map[string]float32            `json:"byStore"`
	TotalSpent     float32                       `json:"totalSpent"`
	TotalPurchases int                           `json:"totalPurchases"`
}

type SaleCreateInput struct {
	StoreID    string `json:"storeId" validate:"required"`
	ClientID   string `json:"clientId" validate:"required"`
	ProductID  string `json:"product"`
	CouponCode string `json:"coupon"`
	Origin     string `json:"origin"`
}

type SaleUpdateInput struct {
	SaleID string           `json:"saleId" validate:"required"`
	Status enums.SaleStatus `json:"status" validate:"required"`
}

type SaleRepository interface {
	GetExpenses(out *SaleGetExpensesOutput, in *SaleGetExpensesInput) error
	Create(out *Sale, in *SaleCreateInput) error
	Update(out *Sale, in *SaleUpdateInput) error
}

//
//
// USECASE
//
//

type SaleCreateUInput struct {
	StoreID   string `json:"storeId" validate:"required"`
	ClientID  string `json:"clientId" validate:"required"`
	ProductID string `json:"productId" validate:"required"`
	Coupon    string `json:"coupon" validate:"required"`
	Origin    string `json:"origin"`
}

type SaleUsecase interface {
	Create(out *Sale, in *SaleCreateUInput) error
	Update(out *Sale, in *SaleUpdateInput) error
}
