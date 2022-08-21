package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/config"
	"github.com/AsaHero/simple-bank/db/driver"
	"github.com/AsaHero/simple-bank/storage/postgres"
)

func TestCreateAccount(t *testing.T) {
	var model *models.CreateAccountRequest
	config := config.Load()

	model = &models.CreateAccountRequest{
		Owner:    "Asadbek",
		Balance:  1000,
		Currency: "UZS",
	}
	db, err := driver.InitPostgresDB(&config)
	assert.NoError(t, err)

	repo := postgres.NewAuthorizationPostgres(db)

	resp, err := repo.CreateAccount(*model)

	assert.NoError(t, err)

	assert.NotEqual(t, 0, resp)
}
