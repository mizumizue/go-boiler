package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func FetchUsers(echo echo.Context) error {
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
