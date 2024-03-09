package repository

import (
	"database/sql"

	"forum/internal/repository/user"
	"forum/internal/types"
)

type Repository struct {
	UserRepo types.UserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: user.NewUserDB(db),
	}
}
