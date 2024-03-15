package usecase

import (
	"context"
	"database/sql"
	"log"

	"github.com/Dzikuri/shopifyx/internal/model"
	"github.com/Dzikuri/shopifyx/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductUseCase struct {
	DB                *sql.DB
	Validate          *validator.Validate
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(db *sql.DB, validate *validator.Validate, productRepository *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		DB:                db,
		ProductRepository: *productRepository,
		Validate:          validate,
	}
}

func (u *ProductUseCase) ProductCreate(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {
    err := u.Validate.Struct(request)
    if err != nil {
        return nil, err
    }

    result, err := u.ProductRepository.Create(ctx, *request)
    if err != nil {
        log.Fatalln(err)
        return nil, echo.ErrBadRequest
    }
    
    return &result, nil
}