package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-bank/storage/postgres"
	"github.com/AsaHero/simple-bank/storage/repo"
)
type Storage struct {
	repo.Authorization
	repo.Transaction
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: postgres.NewAuthorizationPostgres(db),
		Transaction: postgres.NewTransactionPostgres(db),
	}
}
