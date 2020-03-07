package adapter

import (
	"context"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Find(ctx context.Context, userID value_object.UserID) (*entity.User, error)
	Insert(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, userID value_object.UserID) error
}
