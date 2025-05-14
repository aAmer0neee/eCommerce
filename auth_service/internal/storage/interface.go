package storage

import (
	"github.com/aAmer0neee/eCommerce/auth_service/internal/config"
	"github.com/aAmer0neee/eCommerce/auth_service/internal/storage/postgres"
)

type AuthStorage interface {
	GetRecord()
	AddRecord()
	UpdateRecord()
}

func New(cfg *config.Cfg) (AuthStorage, error) { return postgres.NewAuthStorage(cfg) }
