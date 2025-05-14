package service

import (
	"github.com/aAmer0neee/eCommerce/shared/logger"
	"github.com/aAmer0neee/eCommerce/user_service/domain"
	"github.com/aAmer0neee/eCommerce/user_service/internal/storage"
)

type UserService struct {
	storage storage.UserStorage
	log     logger.Logger
}

func newUserService(s storage.UserStorage, l logger.Logger) *UserService {
	return &UserService{
		storage: s,
		log:     l,
	}
}

func (s *UserService) Register(user *domain.RegisterInput) (*domain.RegisterOutput, error) {

	return nil, nil
}
