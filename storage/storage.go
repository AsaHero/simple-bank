package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-bank/storage/postgres"
	"github.com/AsaHero/simple-bank/storage/repo"
)

type Storage interface {
	SimpleBankRepository() repo.SimpleBankRepository
}

type StoragePostgres struct {
	SimpleBankRepo repo.SimpleBankRepository
}

func NewStorage(db *sqlx.DB) Storage {
	return &StoragePostgres{
		SimpleBankRepo: postgres.NewSimpleBankPostgres(db),
	}
}

func (s StoragePostgres) SimpleBankRepository() repo.SimpleBankRepository {
	return s.SimpleBankRepo
}
