package service

import (
	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/storage"
	"github.com/AsaHero/simple-bank/storage/entity"
)

type Service struct {
	Authorization
	Transaction
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthorizationService(storage),
		Transaction: NewTransactionService(storage),
	}
}

type Authorization interface {
	CreateAccount(models.CreateAccountRequest) (uint32, error)
	UpdateAccount(models.UpdateAccountRequest, uint32) error
	GetAccount(uint32) (entity.Account, error)
	DeleteAccount(uint32) error
}

type Transaction interface {

}



