package handler

import "github.com/Dzikuri/shopifyx/internal/usecase"

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}
