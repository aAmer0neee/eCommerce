package postgres

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	Content     string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"not null"`
	Permissions int       `gorm:"default:1"`
}
