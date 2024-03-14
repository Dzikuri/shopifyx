package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserResponse struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserRegisterRequest struct {
	Name     string `json:"name,omitempty" validate:"required,min=5,max=50"`
	Username string `json:"username,omitempty" validate:"required,min=5,max=15"`
	Password string `json:"password,omitempty" validate:"required,max=15"`
}

type UserLoginRequest struct {
	Username string `json:"username,omitempty" validate:"required,min=5,max=15"`
	Password string `json:"password,omitempty" validate:"required,max=15"`
}

type UserAuthResponse struct {
	Username    string `json:"username,omitempty"`
	Name        string `json:"name,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}
