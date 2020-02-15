package service

import (
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) Exists(user *entity.User) bool {
	// TODO 永続化層にアクセスして確認する
	for _, listUser := range inMemoryList {
		if user.Equals(listUser) {
			return true
		}
	}
	return false
}

var (
	un1 = value_object.NewUserName("user1")
)

var inMemoryList = []*entity.User{
	entity.NewUser(un1),
}
