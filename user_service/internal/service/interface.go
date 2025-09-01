package service

import (
	"github.com/aAmer0neee/eCommerce/shared/logger"
	"github.com/aAmer0neee/eCommerce/user_service/domain"
	"github.com/aAmer0neee/eCommerce/user_service/internal/storage"
)

type Service interface {
	Register(user *domain.RegisterInput) (*domain.RegisterOutput, error)
}

func New(s storage.UserStorage, l logger.Logger) Service {
	return newUserService(s, l)
}
