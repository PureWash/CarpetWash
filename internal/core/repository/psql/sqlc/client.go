package sqlc

import (
	"carpet/internal/models"
	"context"
)

const createClientQuery = `
	INSERT INTO orders (
		full_name,
		phone_number,
		latitude,
		longitude
	) VALUES($1, $2, $3, $4)
	RETURNING (
		id,
		full_name,
		phone_number
	)
`

func (q *Queries) CreateClient(ctx context.Context, req models.CreateClientReq) (*models.CreateClientResp, error) {
	var resp models.CreateClientResp

	err := q.db.QueryRow(ctx, createClientQuery,
		req.FullName,
		req.PhoneNumber,
		req.Latitude,
		req.Longitude,
	).
		Scan(&resp.ID, &resp.FullName, &resp.PhoneNumber)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

const updateClientQuery = `
	UPDATE clients 
	SET
		latitude = $1,
		longitude = $2,
		phone_number = $3
	WHERE
		id = $4;
`

func (q *Queries) UpdateClient(ctx context.Context, req models.UpdateClientReq) error {
	_, err := q.db.Exec(ctx, updateClientQuery, 
		req.Latitude, 
		req.Longitude, 
		req.PhoneNumber, 
		req.ID,
	)
	return err
}
