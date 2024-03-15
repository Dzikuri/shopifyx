package model

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type ProductResponse struct {
	Id uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Name string `json:"name,omitempty"`
	Price int `json:"price,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Stock int `json:"stock,omitempty"`
	Condition string `json:"condition,omitempty"`
	Tags pq.StringArray `json:"tags,omitempty"`
	IsPurchasable bool `json:"isPurchasable,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductCreateRequest struct {
	Name string `json:"name,omitempty" validate:"required,min=5,max=60"`
	Price int `json:"price,omitempty" validate:"required,gte=0"`
	ImageUrl string `json:"imageUrl,omitempty" validate:"required,url"`
	Stock int `json:"stock,omitempty" validate:"required,number,gte=0"`
	Condition string `json:"condition,omitempty" validate:"required,oneof=new second"`
	Tags pq.StringArray `json:"tags,omitempty" validate:"required"`
	IsPurchasable bool `json:"isPurchasable,omitempty" validate:"required,boolean"`
    UserId uuid.UUID `json:"userId,omitempty"`
}