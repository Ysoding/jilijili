package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Ysoding/jilijili/app/domain"
	"github.com/Ysoding/jilijili/app/repository"
	"github.com/Ysoding/jilijili/app/repository/dao"
	"github.com/Ysoding/jilijili/app/service"
	"github.com/Ysoding/jilijili/pkg/sqldb"
	"go.uber.org/zap"
)

func UserAdd(log *zap.Logger, cfg sqldb.Config, name, email, password string) error {
	if name == "" || email == "" || password == "" {
		fmt.Println("help: useradd <name> <email> <password>")
		return ErrHelp
	}

	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userSrv := service.NewUserService(repository.NewUserRepository(dao.NewUserDao(db, log)))

	err = userSrv.SignUp(ctx, domain.User{
		Name:     "西门吹雪",
		Email:    "xmchx@test.com",
		Roles:    []string{dao.RoleAdmin},
		Enabled:  true,
		Password: "123456",
	})
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	fmt.Println("user created")
	return nil
}
