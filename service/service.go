package service

import (
	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/storage"
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
}

type Transaction interface {

}



