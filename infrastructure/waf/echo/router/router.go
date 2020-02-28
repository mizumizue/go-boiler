package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trewanek/go-echo-boiler/infrastructure/waf/echo/handler"
)

type Route struct {
	Method      string
	Path        string
	Version     string
	HandlerFunc echo.HandlerFunc
}

var Routes = []*Route{
	{Method: echo.GET, Path: "users", Version: "v1", HandlerFunc: handler.FetchUsers},
	{Method: echo.GET, Path: "users/:user_id", Version: "v1", HandlerFunc: handler.FetchUserByID},
	{Method: echo.POST, Path: "users", Version: "v1", HandlerFunc: handler.CreateUser},
	{Method: echo.PATCH, Path: "users/:user_id", Version: "v1", HandlerFunc: handler.UpdateUser},
	{Method: echo.PUT, Path: "users/:user_id", Version: "v1", HandlerFunc: handler.UpdateUser},
	{Method: echo.DELETE, Path: "users/:user_id", Version: "v1", HandlerFunc: handler.DeleteUser},
}
