--name InsertAddress :exec

INSERT INTO addresses(user_id, latitude, longitude,created_at)
VALUES($1,$2,$3,$4)
RETURNING  id,user_id,latitude,longitude,created_at

--name UpdateAddress :exec
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

--name DeleteAddress :exec
UPDATE
    addresses
SET 
    deleted_at = '0'
WHERE
    id = $1

--name SelectAddress :exec
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
