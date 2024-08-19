package dao

import (
	"context"
	"errors"
	"time"

	"github.com/Ysoding/jilijili/pkg/sqldb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
	RoleAdmin                = "ADMIN"
	RoleUser                 = "USER"
)

// Store
type UserDao interface {
	Create(ctx context.Context, u User) error
	// FindByEmail(ctx context.Context, email string) (User, error)
	// UpdateById(ctx context.Context, entity User) error
	// FindById(ctx context.Context, uid int64) (User, error)
	// FindByPhone(ctx context.Context, phone string) (User, error)
	// FindByWechat(ctx context.Context, openId string) (User, error)
}

type User struct {
	ID          uuid.UUID `db:"user_id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Roles       []string  `db:"roles"`
	Password    string    `db:"password"`
	Enabled     bool      `db:"enabled"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

type SqlxUserDao struct {
	db  sqlx.ExtContext
	log *zap.Logger
}

func NewUserDao(db *sqlx.DB, log *zap.Logger) UserDao {
	return &SqlxUserDao{
		db:  db,
		log: log,
	}
}

func (s *SqlxUserDao) Create(ctx context.Context, u User) error {
	const q = `
	INSERT INTO users
		(user_id, name, email, password, roles, enabled, date_created, date_updated)
	VALUES
		(:user_id, :name, :email, :password, :roles, :enabled, :date_created, :date_updated)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, u); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return ErrUniqueEmail
		}
		return err
	}

	return nil
}
