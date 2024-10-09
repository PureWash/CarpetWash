package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

const InsertServiceQuery = `--name InsertService :exec
INSERT INTO services(tariffs, name, description,price)
VALUES($1,$2,$3,$4)
RETURNING  id,tariffs,name,description,price
`

func (q *Queries) InsertService(ctx context.Context, req *pb.ServiceRequest) (*pb.Service, error) {
	var (
		responses pb.Service
		err       error
	)
	row := q.db.QueryRow(ctx, InsertServiceQuery, req.Tariffs, req.Name, req.Description, req.Price)

	if err = row.Scan(
		&responses.Id,
		&responses.Tariffs,
		&responses.Name,
		&responses.Description,
		&responses.Price,
	); err != nil {
		return nil, err
	}

	return &responses, err
}

const UpdateServiceQuery = `--name UpdateService :exec
UPDATE
    services
SET
    tariffs = $1,
    name = $2,
    description = $3,
    price = $4
WHERE
    id = $5 

RETURNING id,tariffs,name,description,price
`

func (q *Queries) UpdateService(ctx context.Context, req *pb.Service) (*pb.Service, error) {
	var (
		responses pb.Service
		err       error
	)
	row := q.db.QueryRow(ctx, UpdateServiceQuery,
		req.Tariffs,
		req.Name,
		req.Description,
		req.Price,
		req.Id,
	)
	if err = row.Scan(
		&responses.Id,
		&responses.Tariffs,
		&responses.Name,
		&responses.Description,
		&responses.Price,
	); err != nil {
		return nil, err
	}

	return &responses, err
}

const DeleteServiceQuery = `--name DeleteService :exec
DELETE FROM services WHERE id=$1
`

func (q *Queries) DeleteService(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	var (
		err error
	)
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
    price
    
FROM 
    services
WHERE
    id = $1

`

func (q *Queries) SelectService(ctx context.Context, req *pb.PrimaryKey) (*pb.Service, error) {
	var (
		responses pb.Service
		err       error
	)
	row := q.db.QueryRow(ctx, SelectServiceQuery, req.Id)

	if err = row.Scan(
		&responses.Id,
		&responses.Tariffs,
		&responses.Name,
		&responses.Description,
		&responses.Price,
	); err != nil {
		return nil, err
	}

	return &responses, err
}

const SelectServicesQuery = `--name SelectServices
SELECT 
    id,
    tariffs,
    name,
    description,
    price
FROM 
    services
LIMIT $1 OFFSET $2
`

const countSQuery = `--name SelectServices
SELECT 
    COUNT(*)
FROM 
    services
LIMIT $1 OFFSET $2
`

func (q *Queries) SelectServices(ctx context.Context, req *pb.GetListRequest) (*pb.ServicesResponse, error) {
	var (
		err  error
		resp pb.ServicesResponse
	)

	rows, err := q.db.Query(ctx, SelectServicesQuery, req.Limit, req.Page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var responses pb.Service
		if err = rows.Scan(
			&responses.Id,
			&responses.Tariffs,
			&responses.Name,
			&responses.Description,
			&responses.Price,
		); err != nil {
			return nil, err
		}

		resp.Services = append(resp.Services, &responses)
	}

	var count int64
	err = q.db.QueryRow(ctx, countSQuery).Scan(&count)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.ServicesResponse{
		Services:   resp.Services,
		TotalCount: count,
		Limit:      req.Limit,
		Page:       req.Page,
	}, nil
}
