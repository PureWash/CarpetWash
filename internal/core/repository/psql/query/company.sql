--name: InsertCompany :exec
INSERT INTO company(name, description, created_at)
VALUES($1, $2, $3)
RETURNING id, name, description, created_at

--name: UpdateCompany :exec
UPDATE 
    company
SET
    name = $1,
    description = $2
WHERE
    id = $3
AND 
    deleted_at = '1'


--name: DeleteCompany :exec
UPDATE 
    company 
SET    
    deleted_at = '0'
WHERE
    id = $1

--name: SelectCompany :exec
SELECT  
    id,
    name, 
    description, 
    created_at, 
    updated_at
FROM 
    company
WHERE 
    id = $1
AND 
    deleted_at = '1'



