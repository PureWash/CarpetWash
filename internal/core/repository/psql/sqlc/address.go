package sqlc

import (
	pb "carpet/genproto/carpet_service"
	"carpet/internal/configs"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	respons     *pb.Address
	er          error
	createScans sql.NullTime
	updateScans sql.NullTime
)

const InsertAddressQuery = `--name InsertAddress :exec

INSERT INTO addresses(user_id, latitude, longitude,created_at)
VALUES($1,$2,$3,$4)
RETURNING  id,user_id,latitude,longitude,created_at
`

func (q *Queries) InsertAddress(ctx context.Context, req *pb.AddressRequest) (*pb.Address, error) {
	if q.db == nil{
		fmt.Println("+++++++++")
	}
	row := q.db.QueryRow(ctx, InsertAddressQuery, req.UserId, req.Latitude, req.Longitude, time.Now())

	if err = row.Scan(
		&respons.Id,
		&respons.UserId,
		&respons.Latitude,
		&respons.Longitude,
		&createScans,
		&updateScans,
	); err != nil {
		return nil, err
	}
	if createScan.Valid {
		responses.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	if updateScan.Valid {

		responses.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return respons, err
}

const UpdateAddressQuery = `--name UpdateAddress :exec
UPDATE
    addresses
SET
    user_id = $1,
    latitude = $2,
    longitude = $3
WHERE
    id = $4 
AND
    deleted_at = '1'
RETURNING id,user_id,latitude,longitude,updated_at
`

func (q *Queries) UpdateAddress(ctx context.Context, req *pb.Address) (*pb.Address, error) {
	row := q.db.QueryRow(ctx, UpdateAddressQuery, req.UserId, req.Latitude, req.Longitude, time.Now(), req.Id)
	if err = row.Scan(
		&respons.Id,
		&respons.UserId,
		&respons.Latitude,
		&respons.Longitude,
		&updateScans,
	); err != nil {
		return nil, err
	}
	if updateScan.Valid {

		responses.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return respons, err
}

const DeleteAddressQuery = `--name DeleteAddress :exec
UPDATE
    addresses
SET 
    deleted_at = '0'
WHERE
    id = $1
`

func (q *Queries) DeleteAddress(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	_,err = q.db.Exec(ctx, DeleteAddressQuery,req.Id)
	if err != nil {
		return nil, err 
	}
	return &emptypb.Empty{}, nil 
}

const SelectAddressQuery = `--name SelectAddress :exec
SELECT
    id,
    user_id,
    latitude,
    longitude,
    created_at,
    updated_at
FROM 
    addresses
WHERE
    id = $1
AND
    deleted_at = '1'
`
func (q *Queries) SelectAddress(ctx context.Context, req *pb.PrimaryKey) (*pb.Address, error) {
	row := q.db.QueryRow(ctx, SelectAddressQuery,req.Id)

	if err = row.Scan(
		&respons.Id,
		&respons.UserId,
		&respons.Latitude,
		&respons.Longitude,
		&createScans,
		&updateScans,
	); err != nil {
		return nil, err
	}
	if createScan.Valid {
		responses.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	if updateScan.Valid {

		responses.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return respons, err
}
