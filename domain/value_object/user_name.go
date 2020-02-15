package value_object

import "fmt"

type UserName struct {
	userName string
}

func NewUserName(userName string) (*UserName, error) {
	if userName == "" {
		return nil, fmt.Errorf("userName is required")
	}
	if len([]rune(userName)) < 3 {
		return nil, fmt.Errorf("userName is 3 char length required. value: %v", userName)
	}
	return &UserName{userName: userName}, nil
}
