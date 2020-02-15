package entity

import (
	"reflect"
	"testing"

	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

func TestUser_ChangeName(t *testing.T) {
	uid := value_object.NewUserID("example1")
	beforeName := value_object.NewUserName("trewanek")
	afterName := value_object.NewUserName("changedName")

	type fields struct {
		userID   *value_object.UserID
		userName *value_object.UserName
	}
	type args struct {
		name *value_object.UserName
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			"",
			fields{uid, beforeName},
			args{name: afterName},
			&User{uid, afterName},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				userID:   tt.fields.userID,
				userName: tt.fields.userName,
			}

			got, err := func() (got *User, err error) {
				defer func() {
					e := recover()
					if e != nil {
						got = nil
						err = e.(error)
					}
				}()
				return user.ChangeName(tt.args.name), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.ChangeName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.ChangeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Equals(t *testing.T) {
	user1 := NewUser(value_object.NewUserName("trewanek"))
	user2 := NewUser(value_object.NewUserName("trewanek"))
	type fields struct {
		user *User
	}
	type args struct {
		other *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"",
			fields{user1},
			args{user2},
			false,
		},
		{
			"",
			fields{user1},
			args{user1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.user.Equals(tt.args.other); got != tt.want {
				t.Errorf("User.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
