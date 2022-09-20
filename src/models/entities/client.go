package entities

import (
	"time"
)

//
//
// ENTITIES
//
//

type Client struct {
	ID        string    `json:"id" gorm:"column:id;primaryKey;"`
	DiscordID string    `json:"discordId" gorm:"column:discord_id;not null;uniqueIndex;"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null;"`
}

//
//
// REPOSITORY
//
//

type ClientCreateInput struct {
	DiscordID string `json:"discordId" validate:"required"`
}

type ClientRepository interface {
	Get(out *Client, DiscordID string) error
	Create(out *Client, in *ClientCreateInput) error
}

//
//
// USECASE
//
//

type ClientGetProfileOutput struct {
	DiscordID        string  `json:"id"`
	ExpensesCounter  float32 `json:"expensesCounter"`
	PurchasesCounter float32 `json:"purchasesCounter"`
}

type ClientUsecase interface {
	GetProfile(out *ClientGetProfileOutput, DiscordID string) error
	Create(out *Client, in *ClientCreateInput) error
}
