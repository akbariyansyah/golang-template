package service

import (
	"context"
	"task_1/internal/app/param"
	"task_1/internal/domain/user"
)

type UserService interface {
	CreateUser(ctx context.Context, data *param.UserCreate) (*user.User, error)
	ListUsers(ctx context.Context) (user.Users, error)
}

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(ctx context.Context, data *param.UserCreate) (*user.User, error) {
	dataToSave := data.ToUser()
	if err := s.userRepository.Create(ctx, dataToSave); err != nil {
		return nil, err
	}
	return s.userRepository.GetByID(ctx, dataToSave.ID)
}

func (s *userService) ListUsers(ctx context.Context) (user.Users, error) {
	return s.userRepository.List(ctx)
}
