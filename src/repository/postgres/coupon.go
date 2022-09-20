package postgres

import (
	"time"
	"unfinished-api/src/models/entities"

	"gorm.io/gorm"
)

type postgresCouponRepository struct {
	Db *gorm.DB
}

func NewPostgresCouponRepository(Db *gorm.DB) entities.CouponRepository {
	return &postgresCouponRepository{
		Db,
	}
}

func (conn *postgresCouponRepository) Get(out *entities.Coupon, in *entities.CouponGetInput) error {
	tx := conn.Db.
		Where("store_id = ? AND code = ?", in.StoreID, in.Code).
		Take(&out)

	return tx.Error
}

func (conn *postgresCouponRepository) List(out *[]entities.Coupon, in *entities.CouponListInput) error {
	tx := conn.Db.
		Where("store_id = ?", in.StoreID).
		Offset(int(in.Skip)).
		Limit(int(in.Take)).
		Find(&out)

	return tx.Error
}

func (conn *postgresCouponRepository) Create(out *entities.Coupon, in *entities.CouponCreateInput) error {
	out.StoreID = in.StoreID
	out.Code = in.Code
	out.CreatedAt = time.Now()
	out.Code = in.Code
	out.DiscountType = in.DiscountType
	out.DiscountValue = in.DiscountValue
	out.UsesLimit = in.UsesLimit
	out.MinProductValue = in.MinProductValue
	out.ExpiresAt = in.ExpiresAt
	out.CreatedAt = time.Now()

	tx := conn.Db.Create(&out)

	return tx.Error
}

func (conn *postgresCouponRepository) Delete(in *entities.CouponDeleteInput) error {
	tx := conn.Db.
		Where("store_id = ? AND code = ?", in.StoreID, in.Code).
		Update("expires_at", time.Now())

	return tx.Error
}
