package service

import (
	"context"
	"net/mail"

	"github.com/Ysoding/jilijili/app/domain"
	"github.com/Ysoding/jilijili/app/repository"
)

type UserService interface {
}

type userService struct {
	repo repository.UserRepository
}

func (s *userService) Authenticate(ctx context.Context, email mail.Address, password string) (domain.User, error) {
	return domain.User{}, nil
}
