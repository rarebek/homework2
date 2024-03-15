package db

import (
	"FreelancerMarketplace/user-service/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //postgres drivers
)

func ConnectToDB(cfg config.Config) (*sql.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sql.Open("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDb, nil
}
