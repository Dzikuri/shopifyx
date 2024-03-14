package model

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserRegisterRequest struct {
	Username string `json:"username,omitempty" validate:"required,min=5,max=15"`
	Name     string `json:"name,omitempty" validate:"required,min=5,max=50"`
	Password string `json:"password,omitempty" validate:"required,max=15"`
}

type UserAuthResponse struct {
	Username    string `json:"username,omitempty"`
	Name        string `json:"name,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}
