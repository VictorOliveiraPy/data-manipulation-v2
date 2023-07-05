package repository

import (
	"context"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type ClientRawRepository struct {
	Db *pgxpool.Pool
}

func NewClientRawRepository(db *pgxpool.Pool) *ClientRawRepository {
	return &ClientRawRepository{Db: db}
}

func (repository *ClientRawRepository) Create(channelClientRaw <-chan entity.Client, wg *sync.WaitGroup) error {
	batch := &pgx.Batch{}

	for client := range channelClientRaw {
		batch.Queue(
			insertQuery,
			client.ID,
			client.Document,
			client.Private,
			client.Incomplete,
			client.LastPurchaseDate,
			client.TicketAverage,
			client.TicketLastPurchase,
			client.StoreMostFrequent,
			client.StoreLastPurchase,
			client.Status,
			client.CreatedAt,
		)
	}
	br := repository.Db.SendBatch(context.Background(), batch)

	defer wg.Done()
	defer br.Close()
	for range channelClientRaw {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
