package handler

import (
	"net/http"

	"github.com/Dzikuri/shopifyx/internal/model"
	"github.com/Dzikuri/shopifyx/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}

func (h *UserHandler) UserRegister(ctx echo.Context) error {

	request := new(model.UserRegisterRequest)

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(echo.ErrBadRequest.Code, map[string]interface{}{"message": "Bad Request"})
	}

	c := ctx.Request().Context()
	response, err := h.UserUseCase.UserRegister(c, request)

	if err != nil {
		return err
	}

	responseResult := new(model.Response[model.UserAuthResponse])

	responseResult.Data = *response
	responseResult.Message = "User registered successfully"

	return ctx.JSON(http.StatusCreated, responseResult)
}

func (h *UserHandler) UserLogin(ctx echo.Context) error {
	request := new(model.UserLoginRequest)

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(echo.ErrBadRequest.Code, map[string]interface{}{"message": "Bad Request"})
	}

	c := ctx.Request().Context()
	response, err := h.UserUseCase.UserLogin(c, request)

	if err != nil {
		return err
	}

	responseResult := new(model.Response[model.UserAuthResponse])

	responseResult.Data = *response
	responseResult.Message = "User logged successfully"

	return ctx.JSON(http.StatusCreated, responseResult)
}
