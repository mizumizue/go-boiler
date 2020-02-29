package usecase

import (
	"context"

	"github.com/trewanek/go-echo-boiler/application/adapter/repository"
	"github.com/trewanek/go-echo-boiler/domain/entity"
)

type UserInputPort interface {
	FetchUsers(ctx context.Context) error
}

type UserOutputPort interface {
	Present(ctx context.Context, user *entity.User) error
	PresentList(ctx context.Context, users []*entity.User) error
}

type UserUseCase struct {
	userRepo repository.IUserRepository
	output   UserOutputPort
}

func NewUserUseCase(userRepo repository.IUserRepository, output UserOutputPort) UserInputPort {
	return &UserUseCase{
		userRepo: userRepo,
		output:   output,
	}
}

func (u *UserUseCase) FetchUsers(ctx context.Context) error {
	users, err := u.userRepo.Fetch(ctx)
	if err != nil {
		return err
	}
	return u.output.PresentList(ctx, users)
}
