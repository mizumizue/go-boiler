package controller

import (
	"context"

	"github.com/trewanek/go-echo-boiler/application/usecase"
)

type UserController struct {
	input usecase.UserInputPort
}

func NewUserController(input usecase.UserInputPort) *UserController {
	return &UserController{input: input}
}

func (c *UserController) FetchUsers(ctx context.Context) error {
	return c.input.FetchUsers(ctx)
}
