package entity

import (
	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

type IUserFactory interface {
	Create(name value_object.UserName) *User
}

type User struct {
	userID   value_object.UserID
	userName value_object.UserName
}

type IUserDto interface {
	SetUserID(userID string)
	SetUserName(userName string)
}

func (user *User) SetData(dst IUserDto) {
	dst.SetUserID(user.userID.Value())
	dst.SetUserName(user.userName.Value())
}

func NewUser(userID value_object.UserID, userName value_object.UserName) *User {
	return &User{userID: userID, userName: userName}
}

func (user *User) ChangeName(newName value_object.UserName) *User {
	return &User{userID: user.userID, userName: newName}
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return user.userID.Equals(other.userID)
}

func (user *User) UserID() value_object.UserID {
	return user.userID
}
