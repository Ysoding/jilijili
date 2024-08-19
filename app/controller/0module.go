package controller

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("handlers",
		fx.Provide(NewPingController),
		fx.Provide(NewUserController),
	)
}
