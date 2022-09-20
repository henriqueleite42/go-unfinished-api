package entities

import (
	"time"
)

//
//
// ENTITIES
//
//

type Store struct {
	ID             string    `json:"id" gorm:"column:id;primaryKey;"`
	OwnerDiscordID string    `json:"ownerDiscordId" gorm:"column:owner_discord_id;"`
	Name           string    `json:"name" gorm:"column:name;uniqueIndex;"`
	Description    string    `json:"description" gorm:"column:name;"`
	Color          string    `json:"color" gorm:"column:color;"`
	Image          string    `json:"image" gorm:"column:image;"`
	Balance        uint      `json:"balance" gorm:"column:balance;not null;"`
	TotalReceived  uint      `json:"totalReceived" gorm:"column:total_received;not null;"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at;not null;"`
}

//
//
// REPOSITORY
//
//

type GetBalanceOutput struct {
	Balance       uint `json:"balance"`
	TotalReceived uint `json:"totalReceived"`
}

type GetTopOutput struct {
	Stores []Store `json:"stores"`
}

type CreateStoreInput struct {
	OwnerDiscordID string `json:"ownerDiscordId" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description"`
	Color          string `json:"color"`
	Image          string `json:"image"`
}

type StoreRepository interface {
	GetByOwnerDiscordID(out *Store, ownerDiscordID string) error
	GetByName(out *Store, name string) error
	GetBalanceByName(out *GetBalanceOutput, name string) error
	GetTop(out *GetTopOutput) error
	Create(out *Store, in *CreateStoreInput) error
}

//
//
// USECASE
//
//

type StoreUsecase interface {
	GetBalance(out *GetBalanceOutput, name string) error
	GetTop(out *GetTopOutput) error
	Create(out *Store, in *CreateStoreInput) error
}
