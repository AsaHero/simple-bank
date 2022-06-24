package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/storage/entity"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{
		db: db,
	}
}

func (p AuthorizationPostgres) CreateAccount(req models.CreateAccountRequest) (uint32, error) {
	var id uint32

	query := `INSERT INTO account(
			owner,
			balance,
			currency
		) 
		VALUES($1, $2, $3) RETURNING id`

	err := p.db.QueryRow(query, req.Owner, req.Balance, req.Currency).Scan(&id)

	return id, err
}

func (p AuthorizationPostgres) UpdateAccount(req models.CreateAccountRequest) error {
	// db logic
	return nil
}

func (p AuthorizationPostgres) GetAccount(id uint32) (entity.Account, error) {
	// db logic
	return entity.Account{}, nil
}

func (p AuthorizationPostgres) DeleteAccount(id uint32) error {
	// db logic
	return nil
}
