package repository

import (
	"context"

	"github.com/Ysoding/jilijili/app/domain"
	"github.com/Ysoding/jilijili/app/repository/dao"
)

var (
	ErrNotFound              = dao.ErrNotFound
	ErrUniqueEmail           = dao.ErrUniqueEmail
	ErrAuthenticationFailure = dao.ErrAuthenticationFailure
)

type UserRepository interface {
	Create(ctx context.Context, usr domain.User) error
	// FindByEmail(ctx context.Context, email string) (domain.User, error)
	// UpdateNonZeroFields(ctx context.Context, user domain.User) error
	// FindByPhone(ctx context.Context, phone string) (domain.User, error)
	// FindById(ctx context.Context, uid int64) (domain.User, error)
}

type userRepository struct {
	dao dao.UserDao
}

func NewUserRepository(dao dao.UserDao) UserRepository {
	return &userRepository{
		dao: dao,
	}
}

func (r *userRepository) Create(ctx context.Context, usr domain.User) error {
	return r.dao.Create(ctx, r.toEntity(usr))
}

func (repo *userRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Roles:       u.Roles,
		Password:    u.Password,
		Enabled:     u.Enabled,
		DateCreated: u.DateCreated,
		DateUpdated: u.DateUpdated,
	}
}

func (r *userRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Roles:       u.Roles,
		Password:    u.Password,
		Enabled:     u.Enabled,
		DateCreated: u.DateCreated,
		DateUpdated: u.DateUpdated,
	}
}
