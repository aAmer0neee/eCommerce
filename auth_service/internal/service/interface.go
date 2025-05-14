package service

import (
	"github.com/aAmer0neee/eCommerce/auth_service/internal/storage"
	"github.com/aAmer0neee/eCommerce/shared/logger"
)

type Service interface {
}

func New(log logger.Logger, storage storage.AuthStorage) Service {

	return newAuthService(log, storage)
}
