package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			"",
			args{"user1"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUser(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
