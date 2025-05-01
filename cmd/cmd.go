package cmd

import (
	"fmt"
	"log"
	"os"
	"task_1/config"
	"task_1/internal/adapter/repository"
	"task_1/internal/adapter/rest"
	"task_1/internal/app/service"
	"task_1/internal/pkg/logo"
	"task_1/internal/pkg/middleware"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
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
		fmt.Println(err)
		return
	}

	// Register packages
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		log.Fatalf("cannot start postgres database: %v", err)
	}

	// Register Repo
	userRepository := repository.NewUserRepository(db)

	// Register service

	userService := service.NewUserService(userRepository)

	e := echo.New()

	logger, err := middleware.CreateLogger("server.log")
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	e.HideBanner = true
	e.Use(middleware.ErrorHandlingMiddleware)
	e.Use(middleware.LoggerMiddleware(logger))

	// Register handler
	rest.NewUserHandler(e, userService)

	fmt.Println(logo.Logo())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Address)))
}
