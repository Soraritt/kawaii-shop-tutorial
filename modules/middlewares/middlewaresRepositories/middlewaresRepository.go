package middlewaresRepositories

import "github.com/jmoiron/sqlx"

type IMidderwaresRepository interface {
}

type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) IMidderwaresRepository {

	return &middlewaresRepository{
		db: db,
	}
}
