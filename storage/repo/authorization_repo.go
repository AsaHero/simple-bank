package repo

import (
	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/storage/entity"
)

type Authorization interface{
	CreateAccount(models.CreateAccountRequest) (uint32, error)
	UpdateAccount(models.CreateAccountRequest) error
	GetAccount(uint32) (entity.Account, error)
	DeleteAccount(uint32) error
}