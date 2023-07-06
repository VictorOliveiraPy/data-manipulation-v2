package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ClientRawRepository struct {
	Db *pgxpool.Pool
}

func NewClientRawRepository(db *pgxpool.Pool) *ClientRawRepository {
	return &ClientRawRepository{Db: db}
}
