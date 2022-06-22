package service

import "github.com/AsaHero/simple-bank/storage"

type TransactionService struct {
	storage storage.Storage
}

func NewTransactionService(storage *storage.Storage) *TransactionService {
	return &TransactionService{
		storage: *storage,
	}
}