package entity

import (
	"fmt"

	"github.com/trewanek/go-with-ddd-example/domain/derr"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type User struct {
	userID   *value_object.UserID
	userName *value_object.UserName
}

func NewUser(id *value_object.UserID, name *value_object.UserName) *User {
	if id == nil {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("id is nil")))
	}
	if name == nil {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("name is nil")))
	}
	return &User{userID: id, userName: name}
}

func (user *User) ChangeName(name string) *User {
	un := value_object.NewUserName(name)
	return NewUser(user.userID, un)
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return user.userID.Equals(other.userID)
}
