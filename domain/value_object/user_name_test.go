package value_object

import (
	"reflect"
	"testing"
)

func TestNewUserName(t *testing.T) {
	name := "trewanek"
	userName, _ := NewUserName(name)
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserName
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
			nil,
			true,
		},
		{
			"",
			args{"ほげ"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.userName)
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
