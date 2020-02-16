package service

import (
	"context"
	"errors"

	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/aerr"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
)

type UserService struct {
	userRepository adapter.IUserRepository
}

func NewUserService(userRepository adapter.IUserRepository) adapter.IUserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) Exists(ctx context.Context, user *entity.User) (bool, error) {
	var current *entity.User
	var err error
	if current, err = service.userRepository.Find(ctx, user.UserID()); err != nil {
		if errors.As(err, &aerr.ResourceNotFoundErr{}) {
			return false, nil
		}
		return false, err
	}
	return current != nil, nil
}
