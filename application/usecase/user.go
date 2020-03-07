package usecase

import (
	"context"
	"fmt"

	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/aerr"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type UserRegisterUseCase struct {
	userFactory    entity.IUserFactory
	userService    adapter.IUserService
	userRepository adapter.IUserRepository
}

func NewUserRegisterUseCase(
	userFactory entity.IUserFactory,
	userService adapter.IUserService,
	userRepository adapter.IUserRepository,
) *UserRegisterUseCase {
	return &UserRegisterUseCase{
		userFactory:    userFactory,
		userService:    userService,
		userRepository: userRepository,
	}
}

func (u *UserRegisterUseCase) Execute(ctx context.Context, name string) (*entity.User, error) {
	user := u.userFactory.Create(value_object.NewUserName(name))

	exists, err := u.userService.Exists(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user exixts check failed. user: %+v", user)
	}
	if exists {
		return nil, aerr.NewResourceDuplicateErr(fmt.Errorf("user is duplicated. user: %+v", user))
	}

	err = u.userRepository.Insert(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user insert is failed. user: %+v", user)
	}

	return u.userRepository.Find(ctx, user.UserID())
}
