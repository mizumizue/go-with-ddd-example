package main

import (
	"context"
	"log"

	"github.com/trewanek/go-with-ddd-example/interface/presenter"

	"github.com/trewanek/go-with-ddd-example/application/usecase"
	"github.com/trewanek/go-with-ddd-example/domain/service"

	memory2 "github.com/trewanek/go-with-ddd-example/infrastructure/factory/memory"
	memory "github.com/trewanek/go-with-ddd-example/infrastructure/persistence/memory"
)

func main() {
	ctx := context.Background()

	userRepository := memory.NewInMemoryUserRepository()
	userFactory := memory2.NewInMemoryUserFactory()
	userService := service.NewUserService(userRepository)
	userPresenter := presenter.NewUserStdinPresenter()
	use := usecase.NewUserRegisterUseCase(userFactory, userService, userRepository, userPresenter)

	err := use.Execute(ctx, "trewanek")
	if err != nil {
		log.Fatalln(err)
	}

	err = use.Execute(ctx, "trewanek")
	if err != nil {
		log.Fatalln(err)
	}
}
