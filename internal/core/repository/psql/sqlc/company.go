package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/configs"
	"context"
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

const InsertCompanyQuery = `--name: InsertCompany :exec
	INSERT INTO company(name, description, created_at)
	VALUES($1, $2, $3)
	RETURNING id, name, description,created_at,updated_at
`

func (q *Queries) InsertCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error) {
	var (
		response   pb.Company
		err        error
		createScan sql.NullTime
		updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, InsertCompanyQuery, req.Name, req.Description, time.Now())

	if err = row.Scan(
		&response.Id,
		&response.Name,
		&response.Description,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		response.CreatedAt = createScan.Time.Format(configs.Layout)
	}

	if updateScan.Valid {
		response.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}

	return &response, err
}

const UpdateCompanyQuery = `--name: UpdateCompany :exec
	UPDATE 
	    company
	SET
	    name = $1,
	    description = $2,
		updated_at = $3
	WHERE
	    id = $4
	AND 
	    deleted_at = '1'
	RETURNING id, name, description, created_at,updated_at
`

func (q *Queries) UpdateCompany(ctx context.Context, req *pb.Company) (*pb.Company, error) {
	var (
		response   pb.Company
		err        error
		createScan sql.NullTime
		updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, UpdateCompanyQuery,
		req.Name,
		req.Description,
		time.Now(),
		req.Id,
	)

	if err = row.Scan(
		&response.Id,
		&response.Name,
		&response.Description,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		response.CreatedAt = createScan.Time.Format(configs.Layout)
	}

	if updateScan.Valid {
		response.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}

	return &response, nil
}

const DeleteCompanyQuery = `--name: DeleteCompany :exec
	UPDATE 
	    company 
	SET    
	    deleted_at = '0'
	WHERE
	    id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	var (
		err error
	)
	_, err = q.db.Exec(ctx, DeleteCompanyQuery, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

const SelectCompanyQuery = `--name: SelectCompany :exec
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
	    deleted_at = '1'`

func (q *Queries) SelectCompany(ctx context.Context, req *pb.PrimaryKey) (*pb.Company, error) {
	var (
		response   pb.Company
		err        error
		createScan sql.NullTime
		updateScan sql.NullTime
	)
	row := q.db.QueryRow(ctx, SelectCompanyQuery, req.Id)

	if err = row.Scan(
		&response.Id,
		&response.Name,
		&response.Description,
		&createScan,
		&updateScan,
	); err != nil {
		return nil, err
	}

	if createScan.Valid {
		response.CreatedAt = createScan.Time.Format(configs.Layout)
	}

	if updateScan.Valid {
		response.UpdatedAt = updateScan.Time.Format(configs.Layout)
	}

	return &response, nil
}
