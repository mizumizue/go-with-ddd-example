package adapter

import (
	"context"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
)

type IUserService interface {
	Exists(ctx context.Context, user *entity.User) (bool, error)
}
