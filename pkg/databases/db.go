package databases

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/soraritt/kawaii-shop-tutorial/config"
)

func DbConnect(config config.IDbConfig) *sqlx.DB {
	//connect
	db, err := sqlx.Connect("pgx", config.Url())

	if err != nil {
		log.Fatalf("connect to database failed: %v\n", err)
	}

	db.DB.SetMaxOpenConns(config.MaxOpenConns())

	return db
}
