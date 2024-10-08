package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/configs"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// var (
// 	respons     *pb.Address
// 	er          error
// 	createScans sql.NullTime
// 	updateScans sql.NullTime
// )

const InsertAddressQuery = `--name InsertAddress :exec
INSERT INTO addresses(user_id, latitude, longitude,created_at)
VALUES($1,$2,$3,$4)
RETURNING  id,user_id,latitude,longitude,created_at,updated_at
`

func (q *Queries) InsertAddress(ctx context.Context, req *pb.AddressRequest) (*pb.Address, error) {
	if q.db == nil {
		fmt.Println("+++++++++")
	}
	var (
		respons    pb.Address
		err        error
		createScan sql.NullTime
		updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, InsertAddressQuery, req.GetUserId(), req.GetLatitude(), req.GetLongitude(), time.Now())

	if err = row.Scan(
		&respons.Id,
		&respons.UserId,
		&respons.Latitude,
		&respons.Longitude,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}
	if createScan.Valid {
		respons.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	if updateScan.Valid {

		respons.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return &respons, nil
}

const UpdateAddressQuery = `--name UpdateAddress :exec
UPDATE
    addresses
SET
    user_id = $1,
    latitude = $2,
    longitude = $3,
    updated_at=$4
WHERE
    id = $5 
AND
    deleted_at = '1'
RETURNING id,user_id,latitude,longitude,updated_at
`

func (q *Queries) UpdateAddress(ctx context.Context, req *pb.Address) (*pb.Address, error) {
	var (
		response pb.Address
		err      error
		// createScans sql.NullTime
		updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, UpdateAddressQuery, req.UserId, req.Latitude, req.Longitude, time.Now(), req.Id)
	if err = row.Scan(
		&response.Id,
		&response.UserId,
		&response.Latitude,
		&response.Longitude,
		&updateScan,
	); err != nil {
		return nil, err
	}
	if updateScan.Valid {

		response.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return &response, err
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
	var (
		// respons     pb.Address
		err error
		// createScans sql.NullTime
		// updateScans sql.NullTime
	)
	_, err = q.db.Exec(ctx, DeleteAddressQuery, req.Id)
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
	var (
		respons    pb.Address
		err        error
		createScan sql.NullTime
		updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, SelectAddressQuery, req.Id)

	if err = row.Scan(
		&respons.Id,
		&respons.UserId,
		&respons.Latitude,
		&respons.Longitude,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}
	if createScan.Valid {
		respons.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	if updateScan.Valid {

		respons.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}
	return &respons, err
}
