package service

import (
	"github.com/AsaHero/simple-bank/api/models"
	"github.com/AsaHero/simple-bank/storage"
	"github.com/AsaHero/simple-bank/storage/entity"
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

func (s AuthorizationService) UpdateAccount(req models.CreateAccountRequest) error {
	return s.storage.Authorization.UpdateAccount(req)
}
func (s AuthorizationService) GetAccount(id uint32) (entity.Account, error) {
	return s.storage.Authorization.GetAccount(id)
}

func (s AuthorizationService) DeleteAccount(id uint32) error {
	return s.storage.Authorization.DeleteAccount(id)
}