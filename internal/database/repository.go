package database

import (
	"go-away-2024/internal/config"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(cfg *config.Config) *Repository {
	db := Connect(cfg)
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveRequest(request RequestEntity) (id int64, err error) {
	err = r.db.Get(
		&id,
		`insert into request (year, day, part, created_at)
		values ($1, $2, $3, $4) returning id`,
		request.Year,
		request.Day,
		request.Part,
		request.CreatedAt,
	)
	return id, err
}

func (r *Repository) UpdateRequestS3Link(id int64, s3Link string) error {
	_, err := r.db.Exec(
		`update request set s3_link = $2
		where id = $1`,
		id,
		s3Link,
	)
	return err
}

func (r *Repository) SaveResult(rqId int64) error {
	_, err := r.db.Exec(
		`insert into result (request_id)
		values ($1)`,
		rqId,
	)
	return err
}

func (r *Repository) SetResult(result ResultEntity) error {
	_, err := r.db.Exec(
		`update result set
			status = $1,
			result = $2,
			started_at = $3,
			completed_at = $4
		where request_id = $5`,
		result.Status,
		result.Result,
		result.StartedAt,
		result.CompletedAt,
		result.RequestId,
	)
	return err
}

func (r *Repository) GetResult(requestId int64) (res ResultEntity, err error) {
	err = r.db.Get(
		&res,
		`select
			id,
			status,
			result,
			started_at,
			completed_at,
			request_id
		from result
		where request_id = $1`,
		requestId,
	)
	return res, err
}

func (r *Repository) GetRequestWithResult(id int64) (rqRes RequestWithResultEntity, err error) {
	err = r.db.Get(
		&rqRes,
		`select 
			rq.id as request_id,
			rq.year as year,
			rq.day as day,
			rq.part as part,
			rq.created_at as created_at,
			res.started_at as started_at,
			res.completed_at as completed_at,
			res.status as status,
			res.result as result,
			rq.s3_link as s3_link
		from request rq
		left join result res 
			on rq.id = res.request_id
		where rq.id = $1`,
		id,
	)
	return rqRes, err
}
