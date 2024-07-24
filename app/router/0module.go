package router

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers",
		fx.Provide(NewJiliJiliAPIRouter),
	)
}
