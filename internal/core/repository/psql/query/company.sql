--name: InsertCompany :exec
INSERT INTO company(name, description, logo_url, created_at)
VALUES($1, $2, $3, $4)
RETURNING id, name, description, logo_url, created_at

--name: UpdateCompany :exec
	UPDATE 
	    company
	SET
	    name = $1,
	    description = $2,
	    logo_url = $3
		updated_at = $4
	WHERE
	    id = $5
	AND 
	    deleted_at = '1'
	RETURNING id, name, description, logo_url, created_at


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
    logo_url, 
    created_at, 
    updated_at
FROM 
    company
WHERE 
    id = $1
AND 
    deleted_at = '1'

    

