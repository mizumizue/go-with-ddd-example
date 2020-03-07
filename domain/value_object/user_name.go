package value_object

import (
	"fmt"

	"github.com/trewanek/go-with-ddd-example/domain/derr"
)

type UserName struct {
	userName string
}

func NewUserName(userName string) UserName {
	if userName == "" {
		// panic にしているのは値オブジェクトのチェックはあくまでセーフティネットであるので、事前チェックをアプリ側で強制する意味合い
		panic(derr.NewInValidArgumentErr(fmt.Errorf("userName is required")))
	}
	if len([]rune(userName)) < 3 {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("userName is 3 char length required. value: %v", userName)))
	}
	return UserName{userName: userName}
}

func (userName *UserName) Value() string {
	return userName.userName
}
