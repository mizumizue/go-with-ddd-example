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

var userRepository repository.IUserRepository

func CreateUser(ctx context.Context, name string) (*entity.User, error) {
	user := entity.NewUser(value_object.NewUserName(name))

	userService := service.NewUserService()
	if userService.Exists(user) {
		return nil, aerr.NewResorseDuplicateErr(fmt.Errorf("user is duplicated. user: %v", user))
	}

	inserted, err := userRepository.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return inserted, nil
}
