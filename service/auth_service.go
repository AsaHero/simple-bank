package service

import (
	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/storage"
)

type AuthorizationService struct {
	storage storage.Storage
}

func NewAuthorizationService(storage *storage.Storage) *AuthorizationService {
	return &AuthorizationService{
		storage: *storage,
	}
}

func (s AuthorizationService) CreateAccount(req models.CreateAccountRequest) (uint32, error) {
	return s.storage.Authorization.CreateAccount(req)
}