package handler

import (
	"github.com/labstack/echo/v4"
	ierror "github.com/trewanek/go-echo-boiler/interface/error"
	"net/http"
)

func FetchUsers(echo echo.Context) error {
	type req struct {
		Hoge string `json:"hoge" validate:"required"`
	}

	if err := echo.Validate(req{}); err != nil {
		return ierror.NewValidationErr(err)
	}

	return echo.String(http.StatusOK, "Hello, World! FetchUsers Endpoints")
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
