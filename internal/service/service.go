package service

import (
	"forum/internal/repository"
	"forum/internal/service/user"
	"forum/internal/types"
)

type Service struct {
	UserService types.UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
	}
}
