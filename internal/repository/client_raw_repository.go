package repository

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"strconv"
	"sync"
)

type ClientRawRepository struct {
	Db *pgxpool.Pool
}

func NewClientRawRepository(db *pgxpool.Pool) *ClientRawRepository {
	return &ClientRawRepository{Db: db}
}

func (repository *ClientRawRepository) Create(channelClientRaw <-chan entity.ClientRaw, wg *sync.WaitGroup) error {
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

func (repository *ClientRawRepository) GetClients(limit int, status string, channelClientRawRepository chan entity.ClientRaw, wg *sync.WaitGroup) error {
	query := "SELECT id, document, is_private, is_incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store, status FROM raw_client_data WHERE status = $1 LIMIT " + strconv.Itoa(limit)

	rows, err := repository.Db.Query(context.Background(), query, status)
	if err != nil {
		log.Fatal(err)
	}
	defer wg.Done()
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
			return nil
		}
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil
		}
		clients = append(clients, &client)
		channelClientRawRepository <- client
	}
	close(channelClientRawRepository)

	return nil

}
