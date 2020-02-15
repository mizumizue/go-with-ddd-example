package value_object

import (
	"reflect"
	"testing"
)

func TestNewUserID(t *testing.T) {
	id := "hoge"
	uid := NewUserID(id)
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserID
		wantErr bool
	}{
		{
			"",
			args{id},
			uid,
			false,
		},
		{
			"Empty",
			args{""},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := func() (got *UserID, err error) {
				defer func() {
					e := recover()
					if e != nil {
						got = nil
						err = e.(error)
					}
				}()
				return NewUserID(tt.args.userID), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
