package repository

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/jackc/pgx/v5"
)

func (repository *ClientRawRepository) Create(channelClient <-chan entity.Client, channelUpdateClientRaw chan entity.Client) {
	const insertQuery = `INSERT INTO clients (id, document, document_type, private, incomplete, last_purchase_date, ticket_average, ticket_last_purchase, store_most_frequent, store_last_purchase, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id`

	batch := &pgx.Batch{}
	for client := range channelClient {
		batch.Queue(
			insertQuery,
			client.ID,
			client.Document,
			client.DocumentType,
			client.Private,
			client.Incomplete,
			client.LastPurchaseDate,
			client.TicketAverage,
			client.TicketLastPurchase,
			client.StoreMostFrequent,
			client.StoreLastPurchase,
			client.CreatedAt,
			client.UpdatedAt,
		)

	}

	batchResults := repository.Db.SendBatch(context.Background(), batch)
	defer batchResults.Close()
	var clientRaw entity.Client

	for i := 0; i < batch.Len(); i++ {
		row := batchResults.QueryRow()
		err := row.Scan(
			&clientRaw.ID,
		)
		if err != nil {
			fmt.Errorf("scanning create raw batch result")
			continue
		}
		channelUpdateClientRaw <- clientRaw

	}
	close(channelUpdateClientRaw)
}
