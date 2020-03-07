// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/trewanek/go-with-ddd-example/application/adapter"
	"github.com/trewanek/go-with-ddd-example/application/usecase"
	"github.com/trewanek/go-with-ddd-example/domain/service"
	memory2 "github.com/trewanek/go-with-ddd-example/infrastructure/factory/memory"
	memory "github.com/trewanek/go-with-ddd-example/infrastructure/persistence/memory"
)

func InitializeUserRegisterUseCase() *usecase.UserRegisterUseCase {
	wire.Build(
		usecase.NewUserRegisterUseCase,
		service.NewUserService,
		memory.NewInMemoryUserRepository,
		memory2.NewInMemoryUserFactory,
	)
	return nil
}

func InitializeUserService() adapter.IUserService {
	wire.Build(service.NewUserService, memory.NewInMemoryUserRepository)
	return nil
}
