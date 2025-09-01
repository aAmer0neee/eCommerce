package postgres

import (
	"fmt"

	"github.com/aAmer0neee/eCommerce/user_service/domain"
	"github.com/aAmer0neee/eCommerce/user_service/internal/config"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewUserStorage(cfg *config.Cfg) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
		cfg.Postgres.Port,
		cfg.Postgres.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	fmt.Printf("[Repository] [INFO] Open Data Base %s\n", db.Name())

	if cfg.Postgres.Migrate {

		if err := db.AutoMigrate(&User{}); err != nil {
			return nil, err
		}
		fmt.Printf("[Repository][INFO] AutoMigrate")
	}

	return &Postgres{DB: db}, nil
}

func (r *Postgres) CreateUser(user *domain.User) {}
func (r *Postgres) GetUser(Id uuid.UUID)         {}
func (r *Postgres) ModifyUser(user *domain.User) {}
func (r *Postgres) RemoveUser(Id uuid.UUID)      {}
