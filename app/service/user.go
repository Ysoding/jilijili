package service

import (
	"context"
	"net/mail"
	"time"

	"github.com/Ysoding/jilijili/app/domain"
	"github.com/Ysoding/jilijili/app/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUniqueEmail = repository.ErrUniqueEmail
)

type UserService interface {
	SignUp(ctx context.Context, u domain.User) error
	SignIn(ctx context.Context, u domain.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Authenticate(ctx context.Context, email mail.Address, password string) (domain.User, error) {
	return domain.User{}, nil
}

func (s *userService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	u.ID = uuid.New()
	u.Enabled = true
	u.Roles = append(u.Roles, "USER")

	now := time.Now()
	u.DateCreated = now
	u.DateUpdated = now
	return s.repo.Create(ctx, u)
}

func (s *userService) SignIn(ctx context.Context, u domain.User) error {
	return nil
}
