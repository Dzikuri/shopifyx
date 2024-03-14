package config

import (
	"database/sql"

	"github.com/Dzikuri/shopifyx/internal/delivery/handler"
	"github.com/Dzikuri/shopifyx/internal/delivery/routes"
	"github.com/Dzikuri/shopifyx/internal/repository"
	"github.com/Dzikuri/shopifyx/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BootstrapConfig struct {
	DB       *sql.DB
	App      *echo.Echo
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {

	// Setup User
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUseCase(DB, config.Validate, userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	routeConfig := routes.RoutesConfig{
		Echo:        config.App,
		UserHandler: userHandler,
	}

	routeConfig.Setup()
}
