package postgres

import (
	"unfinished-api/src/models/entities"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.Client{},
		&entities.Coupon{},
		&entities.Product{},
		&entities.Sale{},
		&entities.Store{},
	)

	if err != nil {
		panic(err)
	}
}
