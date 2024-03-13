package usecase

import (
	"database/sql"

	"github.com/Dzikuri/shopifyx/internal/repository"
)

type UserUseCase struct {
	DB             *sql.DB
	UserRepository repository.UserRepository
}

func NewUserUseCase(db *sql.DB, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		UserRepository: *userRepository,
	}
}
