package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trewanek/go-echo-boiler/application/usecase"
	"github.com/trewanek/go-echo-boiler/infrastructure/persistence/memory"
	"github.com/trewanek/go-echo-boiler/interface/controller"
	"github.com/trewanek/go-echo-boiler/interface/presenter"
)

func FetchUsers(echo echo.Context) error {
	ctx := echo.Request().Context()
	userRepo := memory.NewUserRepository()
	pre := presenter.NewUserPresenter(echo)
	use := usecase.NewUserUseCase(userRepo, pre)
	ctr := controller.NewUserController(use)
	return ctr.FetchUsers(ctx)
}

func FetchUserByID(echo echo.Context) error {
	return echo.String(http.StatusOK, "Hello, World!")
}

func CreateUser(echo echo.Context) error {
	return echo.String(http.StatusOK, "Hello, World!")
}

func UpdateUser(echo echo.Context) error {
	return echo.String(http.StatusOK, "Hello, World!")
}

func DeleteUser(echo echo.Context) error {
	return echo.String(http.StatusOK, "Hello, World!")
}
