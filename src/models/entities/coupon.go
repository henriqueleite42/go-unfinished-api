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

type Coupon struct {
	StoreID string `json:"storeId" gorm:"primaryKey;column:store_id;not null;"`

	Code          string             `json:"code" gorm:"primaryKey;column:code;not null;"`
	DiscountType  enums.DiscountType `json:"discountType" sql:"type:ENUM('RAW_VALUE', 'PERCENTAGE')" gorm:"column:discount_type;not null;"`
	DiscountValue float32            `json:"discountValue" gorm:"column:discount_value;not null;"`

	UsesCount uint `json:"usesCount" gorm:"column:uses_count;not null;"`
	UsesLimit uint `json:"usesLimit" gorm:"column:uses_limit;not null;"`

	MinProductValue float32 `json:"minProductValue" gorm:"column:min_product_value;"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null;"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"column:expires_at;"`

	Store Store `json:"store"`
}

//
//
// REPOSITORY
//
//

type CouponGetInput struct {
	StoreID string `json:"storeId" validate:"required"`
	Code    string `json:"code" validate:"required"`
}

type CouponListInput struct {
	StoreID string `validate:"required"`
	Take    uint
	Skip    uint
}

type CouponCreateInput struct {
	StoreID string `json:"storeId" validate:"required"`

	Code          string             `json:"code" validate:"required"`
	DiscountType  enums.DiscountType `json:"discountType" validate:"required"`
	DiscountValue float32            `json:"discountValue" validate:"required"`

	UsesLimit uint `json:"usesLimit"`

	MinProductValue float32 `json:"minProductValue"`

	ExpiresAt time.Time `json:"expiresAt"`
}

type CouponDeleteInput struct {
	StoreID string `json:"storeId" validate:"required"`
	Code    string `json:"code" validate:"required"`
}

type CouponRepository interface {
	Get(out *Coupon, in *CouponGetInput) error
	List(out *[]Coupon, in *CouponListInput) error
	Create(out *Coupon, in *CouponCreateInput) error
	Delete(in *CouponDeleteInput) error
}

//
//
// USECASE
//
//

type CouponUsecase interface {
	Get(out *Coupon, in *CouponGetInput) error
	List(out *[]Coupon, in *CouponListInput) error
	Create(out *Coupon, in *CouponCreateInput) error
	Delete(in *CouponDeleteInput) error
}
