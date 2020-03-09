package value

import (
	"reflect"
	"testing"
)

func TestNewUserName(t *testing.T) {
	name := "trewanek"
	userName := NewUserName(name)
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		args    args
		want    UserName
		wantErr bool
	}{
		{
			"",
			args{name},
			userName,
			false,
		},
		{
			"",
			args{"ho"},
			UserName{},
			true,
		},
		{
			"",
			args{"ほげ"},
			UserName{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := func() (got UserName, err error) {
				defer func() {
					e := recover()
					if e != nil {
						err = e.(error)
					}
				}()
				return NewUserName(tt.args.userName), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
