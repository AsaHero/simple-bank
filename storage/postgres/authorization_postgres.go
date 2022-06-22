package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-bank/api/models"
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

	query := fmt.Sprint(
		`INSERT INTO account(
			owner,
			balance,
			currency
		) 
		VALUES($1, $2, $3) RETURNING id`)

	err := p.db.QueryRow(query, req.Owner, req.Balance, req.Currency).Scan(&id)

	return id, err
}
