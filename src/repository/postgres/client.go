package postgres

import (
	"time"

	"unfinished-api/src/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresClientRepository struct {
	Db *gorm.DB
}

func NewPostgresClientRepository(Db *gorm.DB) entities.ClientRepository {
	return &postgresClientRepository{
		Db,
	}
}

func (conn *postgresClientRepository) Get(out *entities.Client, discordID string) error {
	tx := conn.Db.Where("discord_id = ?", discordID).Take(&out)

	return tx.Error
}

func (conn *postgresClientRepository) Create(out *entities.Client, in *entities.ClientCreateInput) error {
	out.ID = uuid.New().String()
	out.DiscordID = in.DiscordID
	out.CreatedAt = time.Now()

	tx := conn.Db.Create(&out)

	return tx.Error
}
