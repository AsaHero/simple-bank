package driver

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/AsaHero/simple-bank/config"
)

func InitPostgresDB(config *config.Config) (*sqlx.DB, error) {
	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.PostgresHost,
		config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDB, config.PostgresSSLMode)
	
	db, err := sqlx.Connect("pgx", postgresInfo)

	if err != nil {
		return nil, err
	}

	return db, nil
}
