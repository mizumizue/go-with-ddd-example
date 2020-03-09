package entity

import (
	"reflect"
	"testing"

	"github.com/trewanek/go-with-ddd-example/domain/value"
)

func TestUser_ChangeName(t *testing.T) {
	type fields struct {
		userID   value.UserID
		userName value.UserName
	}
	type args struct {
		name value.UserName
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
			fields{
				userID:   value.NewUserID("hoge"),
				userName: value.NewUserName("trewanek"),
			},
			args{
				name: value.NewUserName("trewanek2"),
			},
			&User{
				userID:   value.NewUserID("hoge"),
				userName: value.NewUserName("trewanek2"),
			},
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
			fields{
				&User{
					userID:   value.NewUserID("trewanekid"),
					userName: value.NewUserName("trewanek"),
				},
			},
			args{
				&User{
					userID:   value.NewUserID("trewanekid"),
					userName: value.NewUserName("trewanek"),
				},
			},
			true,
		},
		{
			"",
			fields{
				&User{
					userID:   value.NewUserID("trewanekid"),
					userName: value.NewUserName("trewanek"),
				},
			},
			args{
				&User{
					userID:   value.NewUserID("trewanekidhoge"),
					userName: value.NewUserName("trewanek"),
				},
			},
			false,
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
