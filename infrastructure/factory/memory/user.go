package memory

import (
	"strconv"

	"github.com/trewanek/go-with-ddd-example/domain/entity"
	"github.com/trewanek/go-with-ddd-example/domain/value"
)

type InMemoryUserFactory struct {
	seqID int
}

func NewInMemoryUserFactory() entity.IUserFactory {
	return &InMemoryUserFactory{
		seqID: 0,
	}
}

func (factory *InMemoryUserFactory) Create(userName value.UserName) *entity.User {
	factory.seqID++
	userID := value.NewUserID(strconv.Itoa(factory.seqID))
	return entity.NewUser(userID, userName)
}
