package postgres

import (
	"time"

	"unfinished-api/src/models/entities"
	"unfinished-api/src/models/enums"

	"gorm.io/gorm"
)

type postgresSaleRepository struct {
	Db *gorm.DB
}

func NewPostgresSaleRepository(Db *gorm.DB) entities.SaleRepository {
	return &postgresSaleRepository{
		Db,
	}
}

func (conn *postgresSaleRepository) GetExpenses(out *entities.SaleGetExpensesOutput, in *entities.SaleGetExpensesInput) error {
	var startDate time.Time
	var endDate time.Time

	sales := []entities.Sale{}

	tx := conn.Db.Where(
		"client_id = ? AND created_at BETWEEN ? AND ?",
		in.ClientID,
		startDate,
		endDate,
	).Take(&sales)

	if tx.Error != nil {
		return tx.Error
	}

	out.ByStore = make(map[string]float32)
	out.ByType = make(map[enums.ProductType]float32)

	for _, v := range sales {
		productType := enums.ProductType(v.SaleProduct.Type)

		out.ByStore[v.StoreID] += v.FinalPrice
		out.ByType[productType] += v.FinalPrice
		out.TotalSpent += v.FinalPrice
	}
	out.TotalPurchases = len(sales)

	return nil
}

func (conn *postgresSaleRepository) Create(out *entities.Sale, in *entities.SaleCreateInput) error {
	panic("TO DO")
}

func (conn *postgresSaleRepository) Update(out *entities.Sale, in *entities.SaleUpdateInput) error {
	panic("TO DO")
}
