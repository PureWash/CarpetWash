package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/configs"
	"carpet/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

const InsertOrderQuery = `--name: InsertOrder :exec
	INSERT INTO orders (
		client_id,
		service_id,
		area, 
		total_price
	)
	VALUES ($1, $2, $3, $4)
	RETURNING 
		id,
		area,
		total_price,
		created_at
	
`

func (q *Queries) InsertOrder(ctx context.Context, req models.CreateOrderReq) (*models.CreateOrderResp, error) {
	log.Println("go run cmdm")
	var (
		res        models.CreateOrderResp
		err        error
		createScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, InsertOrderQuery,
		req.ClientID,
		req.ServiceId,
		req.Area,
		req.TotalPrice,
	)

	if err = row.Scan(
		&res.ID,
		&res.Area,
		&res.TotalPrice,
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
	    area = $1,
	    status = $2,
		total_price =  $3,
	    updated_at = now()
	WHERE 
	    id = $4
	AND
	    deleted_at = '1'
	RETURNING 
		id, 
		client_id,
		area, 
		total_price, 
		status, 
		updated_at
	
`

func (q *Queries) UpdateOrder(ctx context.Context, req models.UpdateOrderReq) (*models.UpdateOrderResp, error) {
	var (
		res        models.UpdateOrderResp
		err        error
		createScan sql.NullTime
		// resps *pb.OrdersResponse
		// count int64
	)
	row := q.db.QueryRow(ctx, UpdateOrderWithAdmin,
		req.Area,
		req.Status,
		req.TotalPrice,
		req.ID,
	)

	if err = row.Scan(
		&res.ID,
		&res.ClientID,
		&res.Area,
		&res.TotalPrice,
		&res.Status,
		&createScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		res.UpdatedAt = createScan.Time.Format(configs.Layout)
	}

	return &res, nil
}

const DeleteOrderQuery = `--name: DeleteORder :exec
	UPDATE
	    orders
	SET
	    deleted_at = '0'
	WHERE 
	    id = $1;
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
	    o.id,
	    full_name,
		phone_number,
		latitude,
		longitude,
	    name,
		tariffs,
	    area,
	    total_price,
	    status,
	    o.created_at
	FROM    
	    orders AS o
	INNER JOIN clients AS c
		ON o.client_id = c.id AND c.deleted_at = '1'
	INNER JOIN services AS s
		ON o.service_id = s.id
	WHERE
	    o.id = $1
	AND
	    o.deleted_at = '1'
`

func (q *Queries) SelectOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.GetOrderResp, error) {
	var (
		res        pb.GetOrderResp
		err        error
		createScan sql.NullTime
		client     pb.Client
		service    pb.Services
	)
	row := q.db.QueryRow(ctx, SelectOrderQuery, req.Id)

	if err = row.Scan(
		&res.Id,
		&client.FullName,
		&client.PhoneNumber,
		&client.Latitude,
		&client.Longitude,
		&service.Name,
		&service.Tariffs,
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
	res.Client = &client
	res.Service = &service
	return &res, nil
}

var selectOrders = `
	SELECT
		o.id,
		c.id,
		full_name,
		phone_number,
		latitude,
		longitude,
		status
	FROM
		orders AS o
	INNER JOIN clients AS c
		ON o.client_id = c.id
	WHERE
		o.deleted_at = '1' AND
		status = 'READY'
	OFFSET $1 LIMIT $2;
`

var countQuery = `
	SELECT 
		COUNT(*)
	FROM
		orders AS o
	INNER JOIN clients AS c
		ON o.client_id = c.id
	WHERE
		o.deleted_at = '1' AND status = 'READY';
`

func (q *Queries) SelectOrders(ctx context.Context, req *pb.GetListRequest) (*pb.GetOrdersResp, error) {
	var count int32
	err := q.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New(("hech qanday ma'lumot topilmadi"))
	}

	var orders []*pb.Order

	rows, err := q.db.Query(ctx, selectOrders, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order pb.Order
		var client pb.Client

		err := rows.Scan(
			&order.Id, 
			&client.ClientId, 
			&client.FullName, 
			&client.PhoneNumber, 
			&client.Latitude, 
			&client.Longitude, 
			&order.Status,
		)
		if err != nil {
			return nil, err
		}
		order.Client = &client

		orders = append(orders, &order)
	}

	return &pb.GetOrdersResp{
		Orders:     orders,
		TotalCount: count,
		Limit:      int32(req.Limit),
		Offset:     int32(req.Page),
	}, nil
}

var selectOrderFilter = `
SELECT
    o.id,
	c.id,
    c.full_name,
    c.phone_number,
    c.latitude,
    c.longitude,
    o.status
FROM
    orders AS o
INNER JOIN clients AS c
    ON o.client_id = c.id
WHERE
    o.deleted_at = '1'
`

var countsQuery = `
SELECT
    COUNT(*)
FROM
    orders AS o
INNER JOIN clients AS c
    ON o.client_id = c.id
WHERE
    o.deleted_at = '1'
`

func (q *Queries) GetAllOrders(ctx context.Context, req *pb.GetAllOrdersReq) (*pb.GetOrdersResp, error) {
	var (
		filter string
		args   []interface{}
	)

	// FullName filter
	if req.FullName != "" {
		filter += fmt.Sprintf(" AND c.full_name ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", req.FullName))
	}

	// Status filter
	if req.Status != "" {
		filter += fmt.Sprintf(" AND o.status ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", req.Status))
	}

	// OnTime filter (assuming created_at filter)
	if req.OnTime != "" {
		filter += fmt.Sprintf(" AND o.created_at::TEXT ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", req.OnTime))
	}

	// Count total rows
	var count int32
	err := q.db.QueryRow(ctx, countsQuery+filter, args...).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New(("hech qanday ma'lumot topilmadi"))
	}

	// Pagination (LIMIT and OFFSET)
	filter += fmt.Sprintf(" LIMIT %d OFFSET %d", req.Limit, req.Offset)
	// Select orders with the applied filter
	rows, err := q.db.Query(ctx, selectOrderFilter+filter, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Order
	for rows.Next() {
		var order pb.Order
		var client pb.Client
		var status sql.NullString // status uchun sql.NullString dan foydalanamiz

		err := rows.Scan(
			&order.Id, 
			&client.ClientId, 
			&client.FullName, 
			&client.PhoneNumber, 
			&client.Latitude, 
			&client.Longitude, 
			&status,
		)
		if err != nil {
			return nil, err
		}

		// Agar status NULL bo'lsa, default qiymatni beramiz
		if status.Valid {
			order.Status = status.String
		} else {
			order.Status = "" // yoki boshqa default qiymat
		}

		// Assign the client info to the order
		order.Client = &client
		orders = append(orders, &order)
	}

	// Return the response
	return &pb.GetOrdersResp{
		Orders:     orders,
		TotalCount: count,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}, nil
}

var updateOrderStatus = `
	UPDATE orders
	SET
		status = $1,
		updated_at = now()
	WHERE
		id = $2
	RETURNING 
		id
`

func (q *Queries) UpdateOrderStatus(ctx context.Context, req *pb.StatusOrderReq) (*pb.PrimaryKey, error) {
 var ID string
 err := q.db.QueryRow(ctx, updateOrderStatus, req.Status, req.Id).Scan(&ID)
 if err != nil {
  return nil, err
 }
 return &pb.PrimaryKey{
  Id: ID,
 }, nil
}