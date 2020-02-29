package service

import (
	"context"
	"github.com/trewanek/go-echo-boiler/application/adapter/repository"
	"github.com/trewanek/go-echo-boiler/domain/entity"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) Exists(ctx context.Context, user *entity.User) (bool, error) {
	user, err := service.userRepository.FindByID(ctx, user.UserID())
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
