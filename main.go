package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Ysoding/jilijili/app/controller"
	"github.com/Ysoding/jilijili/app/router"
	"github.com/Ysoding/jilijili/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, r *gin.Engine, jilijiliRouter *router.JiliJiliAPIRouter) *http.Server {
	srv := &http.Server{
		Addr:           ":9999",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	rootGroup := r.Group("")
	jilijiliRouter.RegisterPingAPI(rootGroup)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}

func main() {
	app := fx.New(
		fx.Provide(
			NewHTTPServer,
			NewRouter,
			NewLogger,
		),
		service.Module(),
		router.Module(),
		controller.Module(),
		fx.Invoke(func(server *http.Server) {
		}),
	)

	app.Run()
}
