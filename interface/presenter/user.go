package presenter

import (
	"context"
	"os"

	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/response"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
)

type UserStdinPresenter struct {
	*Presenter
}

func NewUserStdinPresenter() adapter.IUserPresenter {
	return &UserStdinPresenter{
		NewPresenter(os.Stdin),
	}
}

func (presenter *UserStdinPresenter) List(ctx context.Context, userList []*entity.User) error {
	panic("implement me")
}

func (presenter *UserStdinPresenter) Detail(ctx context.Context, res *response.UserResponse) error {
	return presenter.outputJSON(res)
}

func (presenter *UserStdinPresenter) Created(ctx context.Context, res *response.UserResponse) error {
	return presenter.outputJSON(res)
}

func (presenter *UserStdinPresenter) Updated(ctx context.Context, res *response.UserResponse) error {
	return presenter.outputJSON(res)
}

func (presenter *UserStdinPresenter) Deleted(ctx context.Context, res *entity.User) error {
	panic("implement me")
}
