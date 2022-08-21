package postgres

import (
	"fmt"
	"strings"

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
	var id uint32 = 0

	query := `INSERT INTO account(
			owner,
			balance,
			currency
		) 
		VALUES($1, $2, $3) RETURNING id`

	err := p.db.QueryRow(query, req.Owner, req.Balance, req.Currency).Scan(&id)

	return id, err
}

func (p AuthorizationPostgres) UpdateAccount(req models.UpdateAccountRequest, id uint32) error {

	sets := make([]string, 0)
	values := make([]interface{}, 0)
	if req.Owner != nil {
		sets = append(sets, "owner=$1")
		values = append(values, req.Owner)

	}

	if req.Balance != nil {
		sets = append(sets, "balance=$2")
		values = append(values, req.Balance)
	}

	if req.Currency != nil {
		sets = append(sets, "currency=$3")
		values = append(values, req.Currency)
	}

	setValues := strings.Join(sets, ", ")

	query := fmt.Sprintf(`UPDATE account SET %s WHERE id = %d`, setValues, id)
	_, err := p.db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error on updating account AuthorizationPostgres -> UpdateAccount: %s", err.Error())
	}

	return nil
}

func (p AuthorizationPostgres) GetAccount(id uint32) (entity.Account, error) {
	account := entity.Account{}
	query := `SELECT * FROM account WHERE id=$1`
	
	err := p.db.Get(&account, query, id)
	
	return account, fmt.Errorf("error on deleting account AuthorizationPostgres -> GetAccount: %s", err)
}

func (p AuthorizationPostgres) DeleteAccount(id uint32) error {
	query := `DELETE FROM account WHERE id = $1`
	_, err := p.db.Exec(query, id)
	
	return fmt.Errorf("error on deleting account AuthorizationPostgres -> DeleteAccount: %s", err)
}
