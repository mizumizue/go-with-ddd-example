package main

import (
	"context"
	"log"

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
	use := usecase.NewUserRegisterUseCase(userFactory, userService, userRepository)
	_, err := use.Execute(ctx, "trewanek1")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = use.Execute(ctx, "trewanek2")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = use.Execute(ctx, "trewanek3")
	if err != nil {
		log.Fatalln(err)
	}

	users, err := userRepository.FindAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range users {
		log.Printf("%++v", v)
	}
}
