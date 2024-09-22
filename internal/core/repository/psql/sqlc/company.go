package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"context"
)

const InsertCompanyQuery = `--name: InsertCompany :exec
insert into company s;ldfads;lfjs`

func (q *Queries) InsertCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error) {
	_, err := q.db.Exec(ctx, InsertCompanyQuery, req)
	if err != nil {
		return nil, err
	}
	return l, err
}


