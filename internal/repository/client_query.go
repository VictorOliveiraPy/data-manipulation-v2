package repository

const insertQuery = `
	INSERT INTO raw_client_data (
		id, document, is_private, is_incomplete, last_purchase_date,
		average_ticket, last_purchase_ticket, most_frequent_store,
		last_purchase_store, status, created_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`
