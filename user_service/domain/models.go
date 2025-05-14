package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id      uuid.UUID
	Email   string
	Name    string
	IsAdmin int

	CreatedAt time.Time
}

type RegisterInput struct {
	Email string
	Name  string
}

type RegisterOutput struct {
	Id string
}
