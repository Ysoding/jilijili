package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Roles       []string
	Password    string
	Enabled     bool
	DateCreated time.Time
	DateUpdated time.Time
}
