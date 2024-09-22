-- name: CreateCompany :one
INSERT INTO  company (id, name, description) 
VALUES ($1, $2, $3)
RETURNING id, name, description;