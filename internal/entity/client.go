package entity

import (
	"github.com/google/uuid"
)

type Client struct {
	ID                 uuid.UUID
	Document           string
	Private            string
	Incomplete         string
	LastPurchaseDate   string
	TicketAverage      string
	TicketLastPurchase string
	StoreMostFrequent  string
	StoreLastPurchase  string
	Status             string
	CreatedAt          string
	UpdatedAt          string
}
