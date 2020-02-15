package entity

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/trewanek/go-with-ddd-example/domain/derr"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type User struct {
	userID   *value_object.UserID
	userName *value_object.UserName
}

func NewUser(id *value_object.UserID, name *value_object.UserName) (*User, error) {
	if id == nil {
		return nil, derr.NewInValidArgumentErr(fmt.Errorf("id is nil"))
	}
	if name == nil {
		return nil, derr.NewInValidArgumentErr(fmt.Errorf("name is nil"))
	}
	return &User{userID: id, userName: name}, nil
}

func CreateUser(name string) (*User, error) {
	uid, err := value_object.NewUserID(uuid.New().String())
	if err != nil {
		return nil, err
	}
	un, err := value_object.NewUserName(name)
	if err != nil {
		return nil, err
	}
	return NewUser(uid, un)
}

func (user *User) ChangeName(name string) (*User, error) {
	un, err := value_object.NewUserName(name)
	if err != nil {
		return nil, err
	}
	return NewUser(user.userID, un)
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return user.userID == other.userID
}
