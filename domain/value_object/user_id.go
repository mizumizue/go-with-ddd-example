package value_object

import "fmt"

type UserID struct {
	userID string
}

func NewUserID(userID string) (*UserID, error) {
	if userID == "" {
		return nil, fmt.Errorf("userID is required")
	}
	return &UserID{userID: userID}, nil
}
