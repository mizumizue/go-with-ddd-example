package service

import (
	"context"
	"errors"

	"github.com/trewanek/go-with-ddd-example/application/adapter/repository"
	"github.com/trewanek/go-with-ddd-example/application/aerr"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) Exists(ctx context.Context, user *entity.User) (bool, error) {
	if _, err := service.userRepository.Find(ctx, user.UserID()); err != nil {
		if errors.As(err, &aerr.ResourceNotFoundErr{}) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
