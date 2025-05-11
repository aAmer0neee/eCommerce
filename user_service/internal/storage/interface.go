package storage

import (
	"github.com/aAmer0neee/eCommerce/user_service/domain"
	"github.com/aAmer0neee/eCommerce/user_service/internal/configModel"
	"github.com/aAmer0neee/eCommerce/user_service/internal/storage/postgres"
	"github.com/google/uuid"
)

type UserStorage interface {
	CreateUser(user *domain.User)
	GetUser(Id uuid.UUID)
	ModifyUser(user *domain.User)
	RemoveUser(Id uuid.UUID)
}

func New(cfg *configModel.Cfg) (UserStorage, error){
	
	return postgres.Connect(cfg)
}
