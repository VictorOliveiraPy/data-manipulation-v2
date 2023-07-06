package repository

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	"time"
)

func (repository *ClientRawRepository) CreateRaw(channelClientRaw <-chan entity.ClientRaw) error {
	const insertQuery = `INSERT INTO raw_client_data (id, document, is_private, is_incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store, status, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

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

	defer br.Close()
	for range channelClientRaw {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *ClientRawRepository) GetClients(limit int, status string, channelClientRawRepository chan entity.ClientRaw) []*entity.ClientRaw {
	query := "SELECT id, document, is_private, is_incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store, status FROM raw_client_data WHERE status = $1 LIMIT " + strconv.Itoa(limit)

	rows, err := repository.Db.Query(context.Background(), query, status)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	clients := make([]*entity.ClientRaw, 0, 1000)

	for rows.Next() {
		var client entity.ClientRaw

		err := rows.Scan(
			&client.ID,
			&client.Document,
			&client.Private,
			&client.Incomplete,
			&client.LastPurchaseDate,
			&client.TicketAverage,
			&client.TicketLastPurchase,
			&client.StoreMostFrequent,
			&client.StoreLastPurchase,
			&client.Status,
		)

		if rows.Err() != nil {
			fmt.Println(rows.Err())
		}
		if err != nil {
			fmt.Println("Error scanning row:", err)
		}
		clients = append(clients, &client)
		channelClientRawRepository <- client
	}
	close(channelClientRawRepository)

	return clients

}

func (repository *ClientRawRepository) UpdateStatusClient(channelClientRawUpdateRepository chan entity.Client, quit chan<- bool) error {
	query := "UPDATE raw_client_data SET status = $2, updated_at = $3 WHERE id = $1"
	batch := &pgx.Batch{}
	status := "concluded"
	updatedAt := time.Now()

	for client := range channelClientRawUpdateRepository {
		batch.Queue(
			query,
			client.ID,
			status,
			updatedAt,
		)
	}

	br := repository.Db.SendBatch(context.Background(), batch)

	defer br.Close()
	quit <- true

	return nil
}
