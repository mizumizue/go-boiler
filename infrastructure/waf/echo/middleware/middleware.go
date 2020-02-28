package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
	MiddlewareFunc echo.MiddlewareFunc
}

var Middlewares = []*Middleware{
	{MiddlewareFunc: middleware.Logger()},
	{MiddlewareFunc: middleware.Recover()},
}
