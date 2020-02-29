package memory

import (
	"context"
	"fmt"

	"github.com/trewanek/go-echo-boiler/application/adapter/repository"
	"github.com/trewanek/go-echo-boiler/domain/entity"
	"github.com/trewanek/go-echo-boiler/domain/value_object"
)

var store = []*entity.User{
	entity.NewUser(value_object.NewUserID(), "user1", 20, "male"),
	entity.NewUser(value_object.NewUserID(), "user2", 30, "female"),
	entity.NewUser(value_object.NewUserID(), "user3", 40, "male"),
	entity.NewUser(value_object.NewUserID(), "user4", 50, "male"),
	entity.NewUser(value_object.NewUserID(), "user5", 60, "female"),
}

type InMemoryUserRepository struct{}

func NewUserRepository() repository.IUserRepository {
	return &InMemoryUserRepository{}
}

func (repo *InMemoryUserRepository) Fetch(ctx context.Context) ([]*entity.User, error) {
	return store, nil
}

func (repo *InMemoryUserRepository) FindByID(ctx context.Context, userID *value_object.UserID) (*entity.User, error) {
	for _, listUser := range store {
		if userID.Equals(listUser.UserID()) {
			return listUser, nil
		}
	}
	return nil, fmt.Errorf("user not found. userID: %v", userID)
}

func (repo *InMemoryUserRepository) Insert(ctx context.Context, user *entity.User) error {
	store = append(store, user)
	return nil
}

func (repo *InMemoryUserRepository) Update(ctx context.Context, user *entity.User) error {
	for index, listUser := range store {
		if user.Equals(listUser) {
			store[index] = user
			return nil
		}
	}
	return fmt.Errorf("user not found. user: %v", user)
}

func (repo *InMemoryUserRepository) Remove(ctx context.Context, userID *value_object.UserID) error {
	newList := make([]*entity.User, 0, len(store)-1)
	for _, listUser := range store {
		if userID.Equals(listUser.UserID()) {
			continue
		}
		newList = append(newList, listUser)
	}
	if len(store) == len(newList) {
		return fmt.Errorf("user not found. userID: %v", userID)
	}
	return nil
}
