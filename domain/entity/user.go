package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	RefreshToken string
	StoresActive int
	CreatedAt    time.Time
}
