--name: InsertOrder :exec
INSERT INTO orders
(user_id, service_id, address_id, area, total_price, status, created_at)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, service_id, address_id, area, total_price, status, created_at


--name: UpdateOrderThisAdmin :exec 
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


--name: UpdateOrderThisUser :exec
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


--name: DeleteORder :exec
UPDATE
    orders
SET
    deleted_at = '0'
WHERE 
    id = $1

--name: SelectOrder :exec
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
    id = $1
AND
    deleted_at = '1'

--name: SelectOrders :many
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

    