--name InsertService :exec
INSERT INTO service(tariffs, name, description,price,created_at)
VALUES($1,$2,$3,$4,$5)
RETURNING  id,tariffs,name,description,price,created_at

--name UpdateService :exec
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

--name DeleteService :exec
UPDATE
    service
SET 
    deleted_at = '0'
WHERE
    id = $1

--name SelectService :exec
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

--name SelectServices
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

