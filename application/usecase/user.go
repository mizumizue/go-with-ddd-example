package usecase

import (
	"github.com/google/uuid"
	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

func CreateUser(name string) *entity.User {
	uid := value_object.NewUserID(uuid.New().String())
	un := value_object.NewUserName(name)
	return entity.NewUser(uid, un)
}
