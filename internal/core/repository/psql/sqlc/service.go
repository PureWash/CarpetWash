package sqlc

import (
	pb "carpet/genproto/carpet_service"
	"carpet/internal/configs"
	"context"
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	responses   *pb.Service
	resp        *pb.ServicesResponse
	errs        error
	createScann sql.NullTime
	updateScann sql.NullTime
)

const InsertServiceQuery = `--name InsertService :exec
INSERT INTO service(tariffs, name, description,price,created_at)
VALUES($1,$2,$3,$4,$5)
RETURNING  id,tariffs,name,description,price,created_at
`

func (q *Queries) InsertService(ctx context.Context, req *pb.ServiceRequest) (*pb.Service, error) {
	row := q.db.QueryRow(ctx, InsertServiceQuery, req.Tariffs, req.Name, req.Description, req.Price, time.Now())

	if err = row.Scan(
		&responses.Id,
		&responses.Tariffs,
		&responses.Name,
		&responses.Description,
		&responses.Price,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}
	if createScan.Valid {
		responses.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	if updateScan.Valid {

		responses.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return responses, err
}

const UpdateServiceQuery = `--name UpdateService :exec
UPDATE
    service
SET
    tariffs = $1,
    name = $2,
    description = $3,
    price = $4
WHERE
    id = $5 
AND
    deleted_at = '1'
RETURNING id,tariffs,name,description,price,updated_at
`

func (q *Queries) UpdateService(ctx context.Context, req *pb.Service) (*pb.Service, error) {
	row := q.db.QueryRow(ctx, UpdateServiceQuery,
		req.Tariffs,
		req.Name,
		req.Description,
		req.Price,
		time.Now(),
		req.Id,
	)
	if err = row.Scan(
		&responses.Id,
		&responses.Tariffs,
		&responses.Name,
		&responses.Description,
		&responses.Price,
		&updateScan,
	); err != nil {
		return nil, err
	}
	if updateScan.Valid {

		responses.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return responses, err
}

const DeleteServiceQuery = `--name DeleteService :exec
UPDATE
    service
SET 
    deleted_at = '0'
WHERE
    id = $1
`

func (q *Queries) DeleteService(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	_, err = q.db.Exec(ctx, DeleteServiceQuery, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

const SelectServiceQuery = `--name SelectService :exec
SELECT
    id,
    tariffs,
    name,
    description,
    price,
    created_at,
    updated_at
FROM 
    service
WHERE
    id = $1
AND
    deleted_at = '1'
`

func (q *Queries) SelectService(ctx context.Context, req *pb.PrimaryKey) (*pb.Service, error) {
	row := q.db.QueryRow(ctx, SelectServiceQuery, req.Id)

	if err = row.Scan(
		&responses.Id,
		&responses.Tariffs,
		&responses.Name,
		&responses.Description,
		&responses.Price,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}
	if createScan.Valid {
		responses.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	if updateScan.Valid {

		responses.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return responses, err
}

const SelectServicesQuery = `--name SelectServices
SELECT 
    id,
    tariffs,
    name,
    description,
    price,
    created_at,
    updated_at
FROM 
    service
WHERE
    id ILIKE $1
OR
    tariffs ILIKE $1
OR
    name ILIKE $1
OR
    description ILIKE $1
OR
    price ILIKE $1
OR 
    created_at ILIKE $1
OR
    updated_at ILIKE $1
AND
    deleted_at not is null
`

func (q *Queries) SelectServices(ctx context.Context, req *pb.GetListRequest) (*pb.ServicesResponse, error) {
	rows, err := q.db.Query(ctx, SelectServicesQuery, req.Limit, req.Page, req.Search)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(
			&responses.Id,
			&responses.Tariffs,
			&responses.Name,
			&responses.Description,
			&responses.Price,
			&createScan,
		); err != nil {
			return nil, err
		}
		if createScan.Valid {
			responses.CreatedAt = createScan.Time.Format(configs.Layout)
		}
		resp.Services = append(resp.Services, responses)
	}
	return &pb.ServicesResponse{
		Services: resp.Services,
	}, nil
}
