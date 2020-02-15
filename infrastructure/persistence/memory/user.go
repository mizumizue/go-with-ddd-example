package memory

import (
	"context"
	"fmt"
	"github.com/trewanek/go-with-ddd-example/application/adapter/repository"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type UserRepository struct {
	list []*entity.User
}

func NewUserRepository(list []*entity.User) repository.IUserRepository {
	return UserRepository{list: list}
}

func (repo UserRepository) Find(ctx context.Context, userID *value_object.UserID) (*entity.User, error) {
	for _, listUser := range repo.list {
		if userID.Equals(listUser.UserID()) {
			return listUser, nil
		}
	}
	return nil, fmt.Errorf("user not found. userID: %v", userID)
}

func (repo UserRepository) Insert(ctx context.Context, user *entity.User) (*entity.User, error) {
	repo.list = append(repo.list, user)
	return user, nil
}

func (repo UserRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	for index, listUser := range repo.list {
		if user.Equals(listUser) {
			repo.list[index] = user
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found. user: %v", user)
}

func (repo UserRepository) Delete(ctx context.Context, userID *value_object.UserID) error {
	newList := make([]*entity.User, 0, len(repo.list)-1)
	for _, listUser := range repo.list {
		if userID.Equals(listUser.UserID()) {
			continue
		}
		newList = append(newList, listUser)
	}
	if len(repo.list) == len(newList) {
		return fmt.Errorf("user not found. userID: %v", userID)
	}
	return nil
}
