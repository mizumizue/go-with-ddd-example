package memory

import (
	"context"
	"fmt"

	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/aerr"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value"
)

var store = map[string]*entity.User{}

type InMemoryUserRepository struct{}

func NewInMemoryUserRepository() adapter.IUserRepository {
	return &InMemoryUserRepository{}
}

func (repo *InMemoryUserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	res := make([]*entity.User, 0, len(store))
	for _, u := range store {
		res = append(res, u)
	}
	return res, nil
}

func (repo *InMemoryUserRepository) Find(ctx context.Context, userID value.UserID) (*entity.User, error) {
	found := store[userID.Value()]
	if found == nil {
		return nil, aerr.NewResourceNotFoundErr(fmt.Errorf("user not found. userID: %v", userID))
	}
	return found, nil
}

func (repo *InMemoryUserRepository) Insert(ctx context.Context, user *entity.User) error {
	store[user.UserID().Value()] = user
	return nil
}

func (repo *InMemoryUserRepository) Update(ctx context.Context, user *entity.User) error {
	found := store[user.UserID().Value()]
	if found == nil {
		return aerr.NewResourceNotFoundErr(fmt.Errorf("user not found. user: %v", user))
	}
	store[user.UserID().Value()] = user
	return nil
}

func (repo *InMemoryUserRepository) Delete(ctx context.Context, userID value.UserID) error {
	found := store[userID.Value()]
	if found == nil {
		return aerr.NewResourceNotFoundErr(fmt.Errorf("user not found. userID: %v", userID))
	}
	delete(store, userID.Value())
	return nil
}
