package storage

import (
	"github.com/aAmer0neee/eCommerce/user_service/domain"
	"github.com/aAmer0neee/eCommerce/user_service/internal/config"
	"github.com/aAmer0neee/eCommerce/user_service/internal/storage/postgres"
	"github.com/google/uuid"
)

type UserStorage interface {
	CreateUser(user *domain.User)
	GetUser(Id uuid.UUID)
	ModifyUser(user *domain.User)
	RemoveUser(Id uuid.UUID)
}

func New(cfg *config.Cfg) (UserStorage, error) {

	return postgres.NewUserStorage(cfg)
}
