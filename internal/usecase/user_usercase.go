package usecase

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Dzikuri/shopifyx/internal/helper"
	"github.com/Dzikuri/shopifyx/internal/model"
	"github.com/Dzikuri/shopifyx/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func (u *UserUseCase) UserLogin(ctx context.Context, request *model.UserLoginRequest) (*model.UserAuthResponse, error) {

	err := u.Validate.Struct(request)
	if err != nil {
		return nil, echo.ErrBadRequest
	}

	result, err := u.UserRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		log.Fatalln(err)
		return nil, echo.ErrBadRequest
	}

	// NOTE check password if matches
	checkPassword := helper.ComparePassword(result.Password, request.Password)
	if checkPassword != nil {
		return nil, errors.New("Password mismatch")
	}

	// NOTE Generate Token

	token, err := helper.JwtGenerateToken(&result)

	if err != nil {
		return nil, echo.ErrBadRequest
	}

	return &model.UserAuthResponse{
		Username: result.Username,
		Name: result.Name,
		AccessToken: token,
	}, nil

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

	// TODO Handling Error for 409 because user is already registered
	if result.Username != "" {
		return nil, errors.New("Username already used")
	}

	// NOTE hash password
	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	requestInsert := new(model.UserRegisterRequest)
	requestInsert.Name = request.Name
	requestInsert.Username = request.Username
	requestInsert.Password = hashedPassword

	result, err = u.UserRepository.Create(ctx, *requestInsert)
	if err != nil {
		return nil, err
	}

	jwtClaims := helper.JwtCustomClaims{
		Name: result.Name,
		Username: result.Username,
		UserId: result.Id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &model.UserAuthResponse{
		Username: result.Username,
		Name: result.Name,
		AccessToken: t,
	}, nil
}
