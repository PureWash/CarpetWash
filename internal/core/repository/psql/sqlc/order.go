package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/configs"
	"carpet/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

const InsertOrderQuery = `--name: InsertOrder :exec
	INSERT INTO orders (
		user_id,
		service_id,
		area, 
		total_price
	)
	VALUES ($1, $2, $3, $4)
	RETURNING (
		id,
		area,
		total_price,
		created_at
	)
`

func (q *Queries) InsertOrder(ctx context.Context, req models.CreateOrderReq) (*models.CreateOrderResp, error) {
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
		time.Now(),
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
	RETURNING (
		id, 
		client_id,
		area, 
		total_price, 
		status, 
		updated_at
	)
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
	)

	if err = row.Scan(
		&res.ID,
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
	    id,
	    full_name,
		phone_number,
		latitude,
		longitude,
	    name,
		tariffs,
	    area,
	    total_price,
	    status,
	    created_at
	FROM    
	    orders AS o
	INNER JOIN clients AS c
		ON o.client_id = c.id AND c.deleted_at IS NULL
	INNER JOIN services AS s
		ON o.service_id = s.id AND s.deleted_at IS NULL
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
		id,
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
		o.deleted_at IS NULL  AND
		status = 'toyyor' AND
		OFFSET = $1 AND
		LIMIT = $2;
`

var countQuery = `
	SELECT 
		COUNT(*)
	FROM
		orders AS o
	INNER JOIN clients AS c
		ON o.client_id = c.id
	WHERE
		o.deleted_at IS NULL AND status = 'toyyor';
`

func (q *Queries) SelectOrders(ctx context.Context, req *pb.GetListRequest) (*pb.GetOrdersResp, error) {
	var orders []*pb.Order

	rows, err := q.db.Query(ctx, selectOrders, (req.Page-1)*req.Limit, req.Limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order pb.Order
		var client pb.Client

		err := rows.Scan(&order.Id, &client.FullName, &client.PhoneNumber, &client.Latitude, &client.Longitude, &order.Status)
		if err != nil {
			return nil, err
		}
		order.Client = &client

		orders = append(orders, &order)
	}

	var count int32
	err = q.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, err
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
		id,
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
		o.deleted_at IS NULL
`

func (q *Queries) GetAllOrders(ctx context.Context, req *pb.GetAllOrdersReq) (*pb.GetOrdersResp, error) {
	var (
		filter string
		args   []interface{}
	)

	if req.FullName != "" {
		filter += fmt.Sprintf(" AND c.full_name ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", req.FullName))
	}
	if req.Status != "" {
		filter += fmt.Sprintf(" AND o.status ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", req.Status))
	}
	if req.OnTime != "" {
		filter += fmt.Sprintf(" AND o.created_at ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", req.OnTime))
	}

	var count int32

	err := q.db.QueryRow(ctx, selectOrderFilter+filter, args...).Scan(&count)
	if err != nil {
		return nil, err
	}

	filter += fmt.Sprintf(" LIMIT %d OFFSET %d", req.Limit, (req.Offset-1)*req.Limit)

	var orders []*pb.Order

	rows, err := q.db.Query(ctx, selectOrderFilter+filter, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order pb.Order
		var client pb.Client

		err := rows.Scan(&order.Id, &client.FullName, &client.PhoneNumber, &client.Latitude, &client.Longitude, &order.Status)
		if err != nil {
			return nil, err
		}
		order.Client = &client
		orders = append(orders, &order)
	}

	return &pb.GetOrdersResp{
		Orders:     orders,
		TotalCount: count,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}, nil
}
