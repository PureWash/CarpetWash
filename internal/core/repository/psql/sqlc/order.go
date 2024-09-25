package sqlc

import (
	pb "carpet/genproto/carpet_service"
	"carpet/internal/configs"
	"context"
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// var (
// 	res   *pb.Order
// 	resps *pb.OrdersResponse
// 	count int64
// )

const InsertOrderQuery = `--name: InsertOrder :exec
	INSERT INTO orders
	(user_id, service_id, area, total_price, status, created_at)
	VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, user_id, service_id, address_id, area, total_price, status, created_at
`

func (q *Queries) InsertOrder(ctx context.Context, req *pb.OrderRequest) (*pb.Order, error) {
	var (
		res        pb.Order
		err        error
		createScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, InsertOrderQuery,
		req.UserId,
		req.ServiceId,
		req.Area,
		req.TotalPrice,
		req.Status,
		time.Now(),
	)

	if err = row.Scan(
		&res.Id,
		&res.UserId,
		&res.ServiceId,
		&res.Area,
		&res.TotalPrice,
		&res.Status,
		&createScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		res.CreatedAt = createScan.Time.Format(configs.Layout)
	}

	return &res, nil
}

const UpdateOrderWithAdmin = `--name: UpdateOrderThisAdmin :exec 
	UPDATE 
	    orders
	SET
	    service_id = $1,
	    area = $2,
	    status = $3,
	    updated_at = $4
	WHERE 
	    id = $5
	AND
	    deleted_at = '1'
	RETURNING id, user_id, service_id, address_id, area, total_price, status, created_at
`

func (q *Queries) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	var (
		res        pb.Order
		err        error
		createScan sql.NullTime
		updateScan sql.NullTime
		// resps *pb.OrdersResponse
		// count int64
	)
	row := q.db.QueryRow(ctx, UpdateOrderWithAdmin,
		req.ServiceId,
		req.Area,
		req.Status,
		time.Now(),
	)

	if err = row.Scan(
		&res.Id,
		&res.UserId,
		&res.ServiceId,
		&res.Area,
		&res.TotalPrice,
		&res.Status,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		res.CreatedAt = createScan.Time.Format(configs.Layout)
	}

	return &res, nil
}

const UpdateOrderWithUser = `--name: UpdateOrderThisUser :exec
	UPDATE
	    orders
	SET 
	    service_id = $1,
	    area = $2
	    updated_at = $3
	WHERE
	    id = $4
	AND
	    deleted_at = '1'
	RETURNING id, user_id, service_id, address_id, area, total_price, status, created_at
`

func (q *Queries) UpdateOrderWithUser(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	var (
		res        pb.Order
		err        error
		createScan sql.NullTime
		// updateScan sql.NullTime
		// resps *pb.OrdersResponse
		// count int64
	)
	row := q.db.QueryRow(ctx, UpdateOrderWithUser,
		req.ServiceId,
		req.Area,
		time.Now(),
	)

	if err = row.Scan(
		&res.Id,
		&res.UserId,
		&res.ServiceId,
		&res.Area,
		&res.TotalPrice,
		&res.Status,
		&createScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		res.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	return &res, nil
}

const DeleteOrderQuery = `--name: DeleteORder :exec
	UPDATE
	    orders
	SET
	    deleted_at = '0'
	WHERE 
	    id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	var (
		err error
	)
	_, err = q.db.Exec(ctx, DeleteCompanyQuery, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

const SelectOrderQuery = `--name: SelectOrder :exec
	SELECT
	    id,
	    user_id,
	    service_id,
	    area,
	    total_price,
	    status,
	    created_at, 
	FROM    
	    orders
	WHERE
	    id = $1
	AND
	    deleted_at = '1'
`

func (q *Queries) SelectOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.Order, error) {
	var (
		res pb.Order
		err error
		createScan sql.NullTime
		// updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, SelectOrderQuery, req.Id)

	if err = row.Scan(
		&res.Id,
		&res.UserId,
		&res.ServiceId,
		&res.Area,
		&res.TotalPrice,
		&res.Status,
		&createScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		res.CreatedAt = createScan.Time.Format(configs.Layout)
	}
	return &res, nil
}

const SelectOrdersQuery = `--name: SelectOrders :many
	SELECT
	    id,
	    user_id,
	    service_id,
	    area,
	    total_price,
	    status,
	    created_at, 
	    updated_at
	FROM
	    orders
	WHERE   
	    id ILIKE $1
	OR
	    user_id ILIKE $1
	OR
	    service_id ILIKE $1
	OR
	    status ILIKE $1
	OR
	    total_price ILIKE $1
	AND
	    deleted_at = '1'
	LIMIT $2 OFFSET $3
`
const OrderCount = `--name: OrderCount :exec
	select 
		COUNT(*) as count
	from 
		orders
	where	
		deleted_at = '1' 
	`

func (q *Queries) SelectOrders(ctx context.Context, req *pb.GetListRequest) (*pb.OrdersResponse, error) {
	var (
		res   pb.Order
		resps pb.OrdersResponse
		count int64
		createScan sql.NullTime
	)

	rows, err := q.db.Query(ctx, SelectOrdersQuery, req.Limit, req.Page, req.Search)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(
			&res.Id,
			&res.UserId,
			&res.ServiceId,
			&res.Area,
			&res.TotalPrice,
			&res.Status,
			&createScan,
		); err != nil {
			return nil, err
		}

		if createScan.Valid {
			res.CreatedAt = createScan.Time.Format(configs.Layout)
		}

		resps.Orders = append(resps.Orders, &res)
	}

	r := q.db.QueryRow(ctx, OrderCount)

	if err = r.Scan(&count); err != nil {
		return nil, err
	}

	return &pb.OrdersResponse{
		Orders: resps.Orders,
		Count:  count,
	}, nil

}
