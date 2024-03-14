package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Dzikuri/shopifyx/internal/helper"
	"github.com/Dzikuri/shopifyx/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (m *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.UserResponse, err error) {
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		// logrus.Error(err)
		log.Println(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			// logrus.Error(errRow)
			log.Println(errRow)

		}
	}()

	result = make([]model.UserResponse, 0)
	for rows.Next() {
		t := model.UserResponse{}
		err = rows.Scan(
			&t.Id,
			&t.Username,
			&t.Name,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			// logrus.Error(err)
			log.Println(err)

			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *UserRepository) getOne(ctx context.Context, query string, args ...interface{}) (res model.UserResponse, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return model.UserResponse{}, err
	}

	defer func() {
		errRow := stmt.Close()
		if errRow != nil {
			// logrus.Error(errRow)
			log.Println(errRow)

		}
	}()

	row := stmt.QueryRowContext(ctx, args...)
	res = model.UserResponse{}

	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.Username,
		&res.Password,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows returned, return an empty result
			return model.UserResponse{}, nil
		}
		return model.UserResponse{}, err
	}

	return
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (res model.UserResponse, err error) {

	query := `SELECT id, username, name, password, created_at, updated_at FROM users WHERE username = $1`
	return r.getOne(ctx, query, username)
}

func (r *UserRepository) Create(ctx context.Context, request model.UserRegisterRequest) (res model.UserResponse, err error) {

	createdAt := time.Now()
	updatedAt := time.Now()

	queryInsert := `INSERT INTO users (username, name, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// _, err = r.DB.Exec(queryInsert, request.Username, request.Name, request.Password, createdAt, updatedAt)

	userId := ""

	err = r.DB.QueryRow(queryInsert, request.Username, request.Name, request.Password, createdAt, updatedAt).Scan(&userId)

	fmt.Println(string(userId))
	if err != nil {
		fmt.Println(err)
		return model.UserResponse{}, err
	} else {
		fmt.Println("value inserted")
	}

	res = model.UserResponse{}
	r.DB.QueryRow("SELECT id, username, name, created_at,updated_at FROM users WHERE id = $1", userId).Scan(
		&res.Id,
		&res.Username,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	return model.UserResponse{
		Id:        helper.GetUUID(userId),
		Username:  res.Username,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
