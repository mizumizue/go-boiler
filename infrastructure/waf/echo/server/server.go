package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/trewanek/go-echo-boiler/infrastructure/validator"
	"github.com/trewanek/go-echo-boiler/infrastructure/waf/echo/handler"
	"github.com/trewanek/go-echo-boiler/infrastructure/waf/echo/middleware"
	"github.com/trewanek/go-echo-boiler/infrastructure/waf/echo/router"
)

const (
	DefaultPort = ":1323"
	Api         = "api"
)

type server struct {
	*echo.Echo
}

func NewServer() *server {
	return &server{
		echo.New(),
	}
}

func (server *server) Run() {
	server.routing()
	server.setMiddleware()
	server.setValidator()
	server.setErrorHandler()
	server.Logger.Fatal(server.Start(DefaultPort))
}

func (server *server) routing() {
	for _, route := range router.Routes {
		path := fmt.Sprintf("/%s/%s/%s", Api, route.Version, route.Path)
		server.Router().Add(route.Method, path, route.HandlerFunc)
	}
}

func (server *server) setMiddleware() {
	for _, m := range middleware.MiddlewareList {
		server.Use(m.MiddlewareFunc)
	}
}

func (server *server) setValidator() {
	server.Validator = validator.New()
}

func (server *server) setErrorHandler() {
	server.HTTPErrorHandler = handler.ErrorHandler
}
