package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id    uuid.UUID `gorm:"primaryKey;type:uuid"`
	Email string    `gorm:"type:varchar(20)"`
	// IpAddress string    `gorm:"type:varchar(50)"`
	TokenHash string    `gorm:"type:varchar(60)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	PairId    uuid.UUID `gorm:"type:uuid"`
}
