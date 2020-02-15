package usecase

import (
	"context"
	"fmt"

	"github.com/trewanek/go-with-ddd-example/application/adapter/repository"
	"github.com/trewanek/go-with-ddd-example/application/aerr"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/service"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type CreateUser struct {
	userService    *service.UserService
	userRepository repository.IUserRepository
}

func NewCreateUser(
	userService *service.UserService,
	userRepository repository.IUserRepository,
) *CreateUser {
	return &CreateUser{
		userService:    userService,
		userRepository: userRepository,
	}
}

func (u *CreateUser) Execute(ctx context.Context, name string) (*entity.User, error) {
	user := entity.NewUser(value_object.NewUserName(name))

	exists, err := u.userService.Exists(ctx, user)
	if err != nil {
		return nil, aerr.NewDatabaseErr(fmt.Errorf("user exixts check failed. user: %v", user))
	}
	if exists {
		return nil, aerr.NewResourceDuplicateErr(fmt.Errorf("user is duplicated. user: %v", user))
	}

	inserted, err := u.userRepository.Insert(ctx, user)
	if err != nil {
		return nil, aerr.NewDatabaseErr(fmt.Errorf("user insert is failed. user: %v", user))
	}

	return inserted, nil
}
