package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
	service2 "github.com/trewanek/go-with-ddd-example/domain/service"
	"github.com/trewanek/go-with-ddd-example/infrastructure/persistence/memory"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	repository := memory.NewUserRepository(nil)
	service := service2.NewUserService(repository)
	usecase := NewUserRegisterUseCase(service, repository)

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO Add Test Cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.Execute(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegisterUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRegisterUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
