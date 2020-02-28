package middleware

import (
	thirdPartyMiddleware "github.com/DeNA/aelog/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
	MiddlewareFunc echo.MiddlewareFunc
}

var MiddlewareList = []*Middleware{
	{MiddlewareFunc: middleware.Logger()},
	{MiddlewareFunc: middleware.Recover()},
	{MiddlewareFunc: echo.WrapMiddleware(thirdPartyMiddleware.AELogger("AELogger"))},
}
