package entity

import (
	"github.com/google/uuid"
	"time"
)

type ClientRaw struct {
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

type Client struct {
	ID                 uuid.UUID
	Document           string
	DocumentType       string
	Private            bool
	Incomplete         bool
	LastPurchaseDate   *time.Time
	TicketAverage      float64
	TicketLastPurchase float64
	StoreMostFrequent  string
	StoreLastPurchase  string
	Status             string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
