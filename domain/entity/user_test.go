package entity

import (
	"reflect"
	"testing"

	"github.com/trewanek/go-with-ddd-example/domain/value_object"
)

func TestNewUser(t *testing.T) {
	uid := value_object.NewUserID("example1")
	un := value_object.NewUserName("trewanek")
	type args struct {
		id   *value_object.UserID
		name *value_object.UserName
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			"",
			args{
				id:   uid,
				name: un,
			},
			&User{
				uid,
				un,
			},
			false,
		},
		{
			"",
			args{
				id:   nil,
				name: un,
			},
			nil,
			true,
		},
		{
			"",
			args{
				id:   uid,
				name: nil,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := func() (got *User, err error) {
				defer func() {
					e := recover()
					if e != nil {
						got = nil
						err = e.(error)
					}
				}()
				return NewUser(tt.args.id, tt.args.name), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			"",
			args{},
			nil,
			true,
		},
		{
			"",
			args{""},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := func() (got *User, err error) {
				defer func() {
					e := recover()
					if e != nil {
						got = nil
						err = e.(error)
					}
				}()
				return CreateUser(tt.args.name), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ChangeName(t *testing.T) {
	uid := value_object.NewUserID("example1")
	un := value_object.NewUserName("trewanek")
	changedName := "changedName"
	changedNameObj := value_object.NewUserName(changedName)

	type fields struct {
		userID   *value_object.UserID
		userName *value_object.UserName
	}
	type args struct {
		name string
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
			fields{uid, un},
			args{name: changedName},
			&User{uid, changedNameObj},
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
	user1 := CreateUser("trewanek")
	user2 := CreateUser("trewanek")
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
