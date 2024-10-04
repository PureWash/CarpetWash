package sqlc

import (
	pb "carpet/genproto/carpet_service"
	"carpet/internal/configs"
	"context"
	"database/sql"
	"encoding/json"
	"log"
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
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id, user_id, service_id,  area, total_price, status, created_at
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
	RETURNING id, user_id, service_id, area, total_price, status, created_at
`

func (q *Queries) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	var (
		res        pb.Order
		err        error
		createScan sql.NullTime
		// resps *pb.OrdersResponse
		// count int64
	)
	row := q.db.QueryRow(ctx, UpdateOrderWithAdmin,
		req.ServiceId,
		req.Area,
		req.Status,
		time.Now(),
		req.Id,
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

//const UpdateOrderWithUser = `--name: UpdateOrderThisUser :exec
//	UPDATE
//	    orders
//	SET
//	    service_id = $1,
//	    area = $2
//	    updated_at = $3
//	WHERE
//	    id = $4
//	AND
//	    deleted_at = '1'
//	RETURNING id, user_id, service_id, area, total_price, status, created_at
//`
//
//func (q *Queries) UpdateOrderWithUser(ctx context.Context, req *pb.Order) (*pb.Order, error) {
//	var (
//		res        pb.Order
//		err        error
//		createScan sql.NullTime
//		// updateScan sql.NullTime
//		// resps *pb.OrdersResponse
//		// count int64
//	)
//	row := q.db.QueryRow(ctx, UpdateOrderWithUser,
//		req.ServiceId,
//		req.Area,
//		time.Now(),
//	)
//
//	if err = row.Scan(
//		&res.Id,
//		&res.UserId,
//		&res.ServiceId,
//		&res.Area,
//		&res.TotalPrice,
//		&res.Status,
//		&createScan,
//	); err != nil {
//		return nil, err
//	}
//
//	if createScan.Valid {
//		res.CreatedAt = createScan.Time.Format(configs.Layout)
//	}
//	return &res, nil
//}

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
	    created_at
	FROM    
	    orders
	WHERE
	    id = $1
	AND
	    deleted_at = '1'
`

func (q *Queries) SelectOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.Order, error) {
	var (
		res        pb.Order
		err        error
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
    o.id,
    (
        SELECT jsonb_agg(
            json_build_object(
                'id', u.id,
                'username', u.username,
                'full_name', u.full_name,        
                'phone_number', u.phone_number    
            )
        )
        FROM users u
        WHERE u.id = o.user_id
    ) AS user_details,
    (
        SELECT jsonb_agg(
            json_build_object(
                'id', s.id,
                'tariffs', s.tariffs,
                'name', s.name,
                'description', s.description,
                'price', s.price
            )
        )
        FROM services s                           
        WHERE s.id = o.service_id
    ) AS service_details,
    o.area,
    o.total_price,
    o.status,
    o.created_at
FROM orders o
WHERE 
    o.deleted_at = '1' AND (
        o.area::text ILIKE $1 OR
        o.total_price::text ILIKE $1 OR
        o.status ILIKE $1 OR
        EXISTS (
            SELECT 1
            FROM services s
            WHERE s.id = o.service_id AND (
                s.tariffs ILIKE $1 OR
                s.name ILIKE $1 OR
                s.description::text ILIKE $1 OR   
                s.price::text ILIKE $1
            )
        )
    )
LIMIT $2 OFFSET $3;                            

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
		res           pb.OrderObject
		resps         pb.OrdersResponse
		count         int64
		createScan    sql.NullTime
		userDetail    json.RawMessage
		serviceDetail json.RawMessage
	)

	rows, err := q.db.Query(ctx, SelectOrdersQuery, req.Search, req.Limit, req.Page)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(
			&res.Id,
			&userDetail,
			&serviceDetail,
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

		if len(userDetail) > 0 {
			var user []*pb.User
			err = json.Unmarshal(userDetail, &user)
			if err != nil {
				log.Println("Unmarshal error:", err)
				return nil, err
			}
			res.UserObject = user
		}
		if len(serviceDetail) > 0 {
			var services []*pb.Service
			err = json.Unmarshal(serviceDetail, &services)
			if err != nil {
				log.Println("Unmarshal error:", err)
				return nil, err
			}
			res.ServiceObject = services
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

//one ishlamadi