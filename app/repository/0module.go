package repository

import (
	"github.com/Ysoding/jilijili/app/repository/dao"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("repository",
		fx.Provide(NewUserRepository),
		fx.Provide(dao.NewUserDao),
	)
}
