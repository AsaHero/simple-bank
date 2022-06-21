package service

import "github.com/AsaHero/simple-bank/storage"

type AuthorizationService struct {
	storage storage.Storage
}

func NewAuthorizationService(storage *storage.Storage) *AuthorizationService {
	return &AuthorizationService{
		storage: *storage,
	}
}