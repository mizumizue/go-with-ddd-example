package adapter

import (
	"context"

	"github.com/trewanek/go-with-ddd-example/application/response"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type IUserPresenter interface {
	List(ctx context.Context, userList []*entity.User) error
	Detail(ctx context.Context, res *response.UserResponse) error
	Created(ctx context.Context, res *response.UserResponse) error
	Updated(ctx context.Context, res *response.UserResponse) error
	Deleted(ctx context.Context, res *entity.User) error
}

type IUserService interface {
	Exists(ctx context.Context, user *entity.User) (bool, error)
}

type IUserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Find(ctx context.Context, userID value_object.UserID) (*entity.User, error)
	Insert(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, userID value_object.UserID) error
}
