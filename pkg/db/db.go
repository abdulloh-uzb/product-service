package db

import (
	"fmt"
	"product-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDB, nil

}
