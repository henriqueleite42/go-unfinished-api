package postgres

import (
	"time"
	"unfinished-api/src/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresStoreRepository struct {
	Db *gorm.DB
}

func NewPostgresStoreRepository(Db *gorm.DB) entities.StoreRepository {
	return &postgresStoreRepository{
		Db,
	}
}

func (conn *postgresStoreRepository) GetByOwnerDiscordID(out *entities.Store, ownerDiscordID string) error {
	tx := conn.Db.Where("owner_discord_id = ?", ownerDiscordID).Take(&out)

	return tx.Error
}

func (conn *postgresStoreRepository) GetByName(out *entities.Store, name string) error {
	tx := conn.Db.Where("name = ?", name).Take(&out)

	return tx.Error
}

func (conn *postgresStoreRepository) GetBalanceByName(out *entities.GetBalanceOutput, name string) error {
	tx := conn.Db.Where("name = ?", name).Take(&out)

	return tx.Error
}

func (conn *postgresStoreRepository) GetTop(out *entities.GetTopOutput) error {
	if out.Stores == nil {
		out.Stores = []entities.Store{}
	}

	tx := conn.Db.Order(`"totalReceived" DESC`).Limit(10).Take(&out.Stores)

	return tx.Error
}

func (conn *postgresStoreRepository) Create(out *entities.Store, in *entities.CreateStoreInput) error {
	out.ID = uuid.New().String()
	out.OwnerDiscordID = in.OwnerDiscordID
	out.Name = in.Name
	out.Description = in.Description
	out.Color = in.Color
	out.Image = in.Image
	out.Balance = 0
	out.TotalReceived = 0
	out.CreatedAt = time.Now()

	tx := conn.Db.Create(&out)

	return tx.Error
}
