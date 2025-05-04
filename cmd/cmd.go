package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"task_1/config"
	"task_1/internal/adapter/repository"
	"task_1/internal/adapter/rest"
	"task_1/internal/app/service"

	"task_1/internal/pkg/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

const (
	path = "config/config"
)

func Run() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSLMODE"),
	)

	config := &config.Configuration{}
	if err := config.SetConfig(path); err != nil {
		log.Println(err)
		return
	}

	logger, err := middleware.CreateLogger("server.log")
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	app := fx.New(
		fx.Supply(databaseURL),
		fx.Supply(config.Address),

		fx.Provide(repository.NewDatabase),

		fx.Provide(
			repository.NewUserRepository,
			service.NewUserService,
		),

		fx.Provide(func() *echo.Echo {
			e := echo.New()
			e.HideBanner = true
			e.Use(middleware.ErrorHandlingMiddleware)
			e.Use(middleware.LoggerMiddleware(logger))
			return e
		}),

		fx.Invoke(func(e *echo.Echo, us service.UserService) {
			rest.NewUserHandler(e, us)
		}),
		fx.Invoke(func(e *echo.Echo) {
			rest.NewHealthHandler(e)
		}),

		fx.Invoke(RegisterServer),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

}
