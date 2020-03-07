package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchUsers(echo echo.Context) error {
	return echo.String(http.StatusOK, "Hello, World!")
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
