package service

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("services",
		fx.Provide(NewTestService),
	)
}
