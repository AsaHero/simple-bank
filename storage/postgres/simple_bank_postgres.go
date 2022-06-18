package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-bank/storage/repo"
)

type SimpleBankPostgres struct {
	db *sqlx.DB
}

func NewSimpleBankPostgres(db *sqlx.DB) repo.SimpleBankRepository {
	return &SimpleBankPostgres{
		db: db,
	}
}

func (p SimpleBankPostgres) CreateAccount() (uint32, error) {
	return 0, nil
}