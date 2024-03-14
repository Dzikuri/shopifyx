package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Dzikuri/shopifyx/internal/helper"
	"github.com/Dzikuri/shopifyx/internal/model"
	"github.com/Dzikuri/shopifyx/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserUseCase struct {
	DB             *sql.DB
	Validate       *validator.Validate
	UserRepository repository.UserRepository
}

func NewUserUseCase(db *sql.DB, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		UserRepository: *userRepository,
		Validate:       validate,
	}
}

func (u *UserUseCase) UserRegister(ctx context.Context, request *model.UserRegisterRequest) (*model.UserAuthResponse, error) {

	err := u.Validate.Struct(request)
	if err != nil {
		return nil, echo.ErrBadRequest
	}
	result, err := u.UserRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		log.Fatalln(err)
		return nil, echo.ErrBadRequest
	}

	fmt.Println(result.Id != uuid.Nil)
	if result.Id != uuid.Nil {
		return nil, errors.New("Username already used")
	}

	helper.LogPretty(result)

	return &model.UserAuthResponse{}, nil
}
