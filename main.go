package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/Ysoding/jilijili/app/controller"
	"github.com/Ysoding/jilijili/app/repository"
	"github.com/Ysoding/jilijili/app/service"
	"github.com/Ysoding/jilijili/app/web"
	"github.com/Ysoding/jilijili/pkg/sqldb"
	"github.com/ardanlabs/conf/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var build = "develop"

type Config struct {
	Version conf.Version
	Web     struct {
		ReadTimeout        time.Duration `conf:"default:5s"`
		WriteTimeout       time.Duration `conf:"default:10s"`
		IdleTimeout        time.Duration `conf:"default:120s"`
		ShutdownTimeout    time.Duration `conf:"default:20s"`
		APIHost            string        `conf:"default:0.0.0.0:9000"`
		CORSAllowedOrigins []string      `conf:"default:*"`
	}
	DB struct {
		User         string `conf:"default:postgres"`
		Password     string `conf:"default:postgres,mask"`
		Host         string `conf:"default:database-service"`
		Name         string `conf:"default:postgres"`
		MaxIdleConns int    `conf:"default:0"`
		MaxOpenConns int    `conf:"default:0"`
		DisableTLS   bool   `conf:"default:true"`
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	_, err := conf.Parse("JILI", cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			return nil, fmt.Errorf("parsing config: %w", err)
		}
	}
	return cfg, nil
}

func NewDatabase(cfg *Config, lc fx.Lifecycle, log *zap.Logger) (*sqlx.DB, error) {
	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return nil, fmt.Errorf("connecting to db: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return sqldb.StatusCheck(ctx, db)
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}

func NewHTTPServer(lc fx.Lifecycle, r *gin.Engine, log *zap.Logger, jilijiliRouter *web.JiliJiliAPIRouter, cfg *Config) *http.Server {
	log.Info("startup", zap.Int("GOMAXPROCS", runtime.GOMAXPROCS(0)))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Web.CORSAllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}))

	srv := &http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      r,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
	}

	jilijiliRouter.RegisterRoutes(r)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("starting service", zap.String("version", cfg.Version.Build))
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer log.Info("service down")
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
	zap.ReplaceGlobals(logger)
	return logger
}

func main() {
	app := fx.New(
		fx.Provide(
			NewHTTPServer,
			NewRouter,
			NewLogger,
			NewConfig,
			NewDatabase,
		),
		web.Module(),
		controller.Module(),
		service.Module(),
		repository.Module(),
		fx.Invoke(func(server *http.Server) {
		}),
	)

	app.Run()
}
