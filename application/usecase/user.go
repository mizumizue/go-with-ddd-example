package usecase

import (
	"context"
	"fmt"

	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/aerr"
	"github.com/trewanek/go-with-ddd-example/application/response"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value"
)

type RegisterUserUseCase struct {
	userFactory    entity.IUserFactory
	userService    adapter.IUserService
	userRepository adapter.IUserRepository
	userPresenter  adapter.IUserPresenter
}

func NewUserRegisterUseCase(
	userFactory entity.IUserFactory,
	userService adapter.IUserService,
	userRepository adapter.IUserRepository,
	userPresenter adapter.IUserPresenter,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userFactory:    userFactory,
		userService:    userService,
		userRepository: userRepository,
		userPresenter:  userPresenter,
	}
}

func (u *RegisterUserUseCase) Execute(ctx context.Context, name string) error {
	user := u.userFactory.Create(value.NewUserName(name))

	exists, err := u.userService.Exists(ctx, user)
	if err != nil {
		return fmt.Errorf("user exixts check failed. user: %+v", user)
	}
	if exists {
		return aerr.NewResourceDuplicateErr(fmt.Errorf("user is duplicated. user: %+v", user))
	}

	err = u.userRepository.Insert(ctx, user)
	if err != nil {
		return fmt.Errorf("user insert is failed. user: %+v", user)
	}

	found, err := u.userRepository.Find(ctx, user.UserID())
	if err != nil {
		return fmt.Errorf("find user by id is failed. user: %+v", user)
	}

	res := new(response.UserResponse)
	found.SetData(res)
	return u.userPresenter.Created(ctx, res)
}
