package domain

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	// Name         Name
	Email mail.Address
	// Roles        []Role
	PasswordHash []byte
	Department   string
	Enabled      bool
	DateCreated  time.Time
	DateUpdated  time.Time
}
