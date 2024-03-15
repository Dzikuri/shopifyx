package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Dzikuri/shopifyx/internal/helper"
	"github.com/Dzikuri/shopifyx/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(ctx context.Context, request model.ProductCreateRequest) (res model.ProductResponse, err error) {
    createdAt := time.Now()
	updatedAt := time.Now()

	queryInsert := `INSERT INTO products (user_id, name, price, image_url, stock, condition, tags, is_purchasable, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	productId := ""

    err = r.DB.QueryRow(queryInsert, request.UserId,request.Name, request.Price,request.ImageUrl,request.Stock,request.Condition,request.Tags,request.IsPurchasable,createdAt,updatedAt).Scan(&productId)

	if err != nil {
		fmt.Println(err)
		return model.ProductResponse{}, err
	} else {
		fmt.Println("value inserted")
	}

    res = model.ProductResponse{}
	r.DB.QueryRow("SELECT id, user_id, name, price, image_url, stock, condition, tags, is_purchasable, created_at, updated_at FROM products WHERE id = $1", productId).Scan(
        &res.Id,
        &res.UserId,
        &res.Name,
        &res.Price,
        &res.ImageUrl,
        &res.Stock,
        &res.Condition,
        &res.Tags,
        &res.IsPurchasable,
        &res.UpdatedAt,
        &res.CreatedAt,
    )

    return model.ProductResponse{
        Id: helper.GetUUID(productId),
        UserId: helper.GetUUID(res.UserId.String()),
        Name: res.Name,
        Price: res.Price,
        ImageUrl: res.ImageUrl,
        Stock: res.Stock,
        Condition: res.Condition,
        Tags: res.Tags,
        IsPurchasable: res.IsPurchasable,
        UpdatedAt: res.UpdatedAt,
        CreatedAt: res.CreatedAt,
    }, nil
}