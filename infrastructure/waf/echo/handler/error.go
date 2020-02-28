package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trewanek/go-echo-boiler/infrastructure/http/status"
	"github.com/trewanek/go-echo-boiler/infrastructure/validator"
	ierror "github.com/trewanek/go-echo-boiler/interface/error"
	"github.com/trewanek/go-echo-boiler/interface/response"
)

func ErrorHandler(err error, echo echo.Context) {
	if errors.As(err, &ierror.ValidationErr{}) {
		v := echo.Echo().Validator.(*validator.Validator)
		errorResponse(echo, http.StatusBadRequest, status.Invalid, v.ValidationStrings(errors.Unwrap(err))...)
		return
	}
	errorResponse(echo, http.StatusInternalServerError, status.Unknown, err.Error())
}

func errorResponse(echo echo.Context, code int, status string, errStrings ...string) {
	err := echo.JSON(code, response.HttpResponse{
		Code:   code,
		Status: status,
		Errors: errStrings,
	})
	if err != nil {
		echo.Logger().Errorf("err: ", err)
	}
}
