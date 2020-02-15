package entity

import (
	"github.com/google/uuid"
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type User struct {
	userID   *value_object.UserID
	userName *value_object.UserName
}

func NewUser(userName *value_object.UserName) *User {
	return &User{userID: value_object.NewUserID(uuid.New().String()), userName: userName}
}

func (user *User) ChangeName(newName *value_object.UserName) *User {
	return &User{userID: user.userID, userName: newName}
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return user.userID.Equals(other.userID)
}
