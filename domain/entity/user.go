package entity

import (
	"github.com/trewanek/go-with-ddd-example/domain/value"
)

type IUserFactory interface {
	Create(name value.UserName) *User
}

type User struct {
	userID   value.UserID
	userName value.UserName
}

type IUserDto interface {
	SetUserID(userID string)
	SetUserName(userName string)
}

func (user *User) SetData(dst IUserDto) {
	dst.SetUserID(user.userID.Value())
	dst.SetUserName(user.userName.Value())
}

func NewUser(userID value.UserID, userName value.UserName) *User {
	return &User{userID: userID, userName: userName}
}

func (user *User) ChangeName(newName value.UserName) *User {
	return &User{userID: user.userID, userName: newName}
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return user.userID.Equals(other.userID)
}

func (user *User) UserID() value.UserID {
	return user.userID
}
