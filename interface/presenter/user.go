package presenter

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trewanek/go-echo-boiler/application/usecase"
	"github.com/trewanek/go-echo-boiler/domain/entity"
)

type UserPresenter struct {
	e echo.Context
}

type UserDto struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
}

func NewUserDto(entity *entity.User) *UserDto {
	return &UserDto{
		UserID:   entity.UserID().Value(),
		UserName: entity.UserName(),
		Age:      entity.Age(),
		Sex:      entity.Sex(),
	}
}

func NewUserDtoList(entities []*entity.User) []*UserDto {
	list := make([]*UserDto, 0, len(entities))
	for _, e := range entities {
		list = append(list, NewUserDto(e))
	}
	return list
}

func NewUserPresenter(echo echo.Context) usecase.UserOutputPort {
	return &UserPresenter{e: echo}
}

func (p *UserPresenter) Present(ctx context.Context, user *entity.User) error {
	return p.e.JSON(http.StatusOK, NewUserDto(user))
}

func (p *UserPresenter) PresentList(ctx context.Context, users []*entity.User) error {
	return p.e.JSON(http.StatusOK, NewUserDtoList(users))
}
