package service

import (
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
	"reflect"
	"testing"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
)

func TestNewUserService(t *testing.T) {
	tests := []struct {
		name string
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Exists(t *testing.T) {
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		service *UserService
		args    args
		want    bool
	}{
		{
			"",
			NewUserService(),
			args{entity.NewUser(
				value_object.NewUserID("1"),
				value_object.NewUserName("user1"),
			)},
			true,
		},
		{
			"",
			NewUserService(),
			args{entity.NewUser(
				value_object.NewUserID("2"),
				value_object.NewUserName("user2"),
			)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.Exists(tt.args.user); got != tt.want {
				t.Errorf("UserService.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
