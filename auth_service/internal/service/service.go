package service

import (
	"github.com/aAmer0neee/eCommerce/auth_service/internal/storage"
	"github.com/aAmer0neee/eCommerce/shared/logger"
)

type AuthService struct {
	log     logger.Logger
	storage storage.AuthStorage
}

func newAuthService(log logger.Logger, storage storage.AuthStorage) *AuthService {
	return &AuthService{
		log:     log,
		storage: storage,
	}
}
