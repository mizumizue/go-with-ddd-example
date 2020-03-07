package controller

import (
	"context"

	"github.com/trewanek/go-with-ddd-example/application/usecase"
)

type UserController struct {
	registerUserUseCase usecase.RegisterUserUseCase

	// TODO impl
	//fetchUserListUseCase usecase.FetchUserListUseCase
	//getUserDetailUseCase usecase.GetUserDetailUseCase
	//updateByIDUseCase    usecase.UserFindByIDUseCase
	//withdrawUserUseCase  usecase.WithdrawUserUseCase
}

func NewUserController(registerUseCase usecase.RegisterUserUseCase) *UserController {
	return &UserController{registerUserUseCase: registerUseCase}
}

func (ctr *UserController) Register(ctx context.Context, userName string) error {
	return ctr.registerUserUseCase.Execute(ctx, userName)
}
