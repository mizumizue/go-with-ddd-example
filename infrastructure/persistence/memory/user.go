package memory

import (
	"context"
	"fmt"

	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type InMemoryUserRepository struct {
	store []*entity.User
}

func NewUserRepository(store []*entity.User) adapter.IUserRepository {
	return &InMemoryUserRepository{store: store}
}

func (repo *InMemoryUserRepository) Find(ctx context.Context, userID *value_object.UserID) (*entity.User, error) {
	for _, listUser := range repo.store {
		if userID.Equals(listUser.UserID()) {
			return listUser, nil
		}
	}
	return nil, fmt.Errorf("user not found. userID: %v", userID)
}

func (repo *InMemoryUserRepository) Insert(ctx context.Context, user *entity.User) (*entity.User, error) {
	repo.store = append(repo.store, user)
	return user, nil
}

func (repo *InMemoryUserRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	for index, listUser := range repo.store {
		if user.Equals(listUser) {
			repo.store[index] = user
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found. user: %v", user)
}

func (repo *InMemoryUserRepository) Delete(ctx context.Context, userID *value_object.UserID) error {
	newList := make([]*entity.User, 0, len(repo.store)-1)
	for _, listUser := range repo.store {
		if userID.Equals(listUser.UserID()) {
			continue
		}
		newList = append(newList, listUser)
	}
	if len(repo.store) == len(newList) {
		return fmt.Errorf("user not found. userID: %v", userID)
	}
	return nil
}
