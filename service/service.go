package service

import "github.com/AsaHero/simple-bank/storage"

type Authorization interface {
	
}

type Transaction interface {

}

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


