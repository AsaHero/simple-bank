package repo

import "github.com/AsaHero/simple-bank/api/models"

type Authorization interface{
	CreateAccount(models.CreateAccountRequest) (uint32, error)
}