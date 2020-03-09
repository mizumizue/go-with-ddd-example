package value

import (
	"fmt"

	"github.com/trewanek/go-with-ddd-example/domain/derr"
)

type UserID struct {
	userID string
}

func NewUserID(userID string) UserID {
	if userID == "" {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("userID is required")))
	}
	return UserID{userID: userID}
}

func (id UserID) Value() string {
	return id.userID
}

func (id UserID) Equals(other UserID) bool {
	return id.userID == other.userID
}
