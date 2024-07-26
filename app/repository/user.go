package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Ysoding/jilijili/app/domain"
	"github.com/Ysoding/jilijili/pkg/sqldb"
	"github.com/Ysoding/jilijili/pkg/sqldb/dbarray"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

type UserRepository interface {
	Create(ctx context.Context, usr domain.User) error
}

type userRepository struct {
	log *zap.Logger
	db  sqlx.ExtContext
}

func NewUserRepository(db *sqlx.DB, log *zap.Logger) UserRepository {
	return &userRepository{
		log: log,
		db:  db,
	}
}

func (r *userRepository) Create(ctx context.Context, usr domain.User) error {
	const q = `
	INSERT INTO users
		(user_id, name, email, password_hash, roles, enabled, date_created, date_updated)
	VALUES
		(:user_id, :name, :email, :password_hash, :roles, :enabled, :date_created, :date_updated)`

	if err := sqldb.NamedExecContext(ctx, r.log, r.db, q, toDBUser(usr)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", ErrUniqueEmail)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Store
type user struct {
	ID           uuid.UUID      `db:"user_id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Roles        dbarray.String `db:"roles"`
	PasswordHash []byte         `db:"password_hash"`
	Department   sql.NullString `db:"department"`
	Enabled      bool           `db:"enabled"`
	DateCreated  time.Time      `db:"date_created"`
	DateUpdated  time.Time      `db:"date_updated"`
}

func toDomainUser(db user) (domain.User, error) {
	bus := domain.User{
		ID:           db.ID,
		PasswordHash: db.PasswordHash,
		Enabled:      db.Enabled,
		Department:   db.Department.String,
		DateCreated:  db.DateCreated.In(time.Local),
		DateUpdated:  db.DateUpdated.In(time.Local),
	}
	return bus, nil
}

func toDBUser(bus domain.User) user {
	return user{
		ID:           bus.ID,
		Email:        bus.Email.Address,
		PasswordHash: bus.PasswordHash,
		Department: sql.NullString{
			String: bus.Department,
			Valid:  bus.Department != "",
		},
		Enabled:     bus.Enabled,
		DateCreated: bus.DateCreated.UTC(),
		DateUpdated: bus.DateUpdated.UTC(),
	}
}
