package database

func SaveRequest(request RequestEntity) (id int64, err error) {
	err = DB.Get(
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
