package repository

import (
	"context"

	"github.com/trewanek/go-echo-boiler/domain/entity"
	"github.com/trewanek/go-echo-boiler/domain/value_object"
)

type IUserRepository interface {
	Fetch(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, userID *value_object.UserID) (*entity.User, error)
	Insert(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Remove(ctx context.Context, userID *value_object.UserID) error
}
