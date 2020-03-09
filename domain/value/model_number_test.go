package value

import (
	"reflect"
	"testing"
)

func TestNewModelNumber(t *testing.T) {
	type args struct {
		productCode string
		branch      string
		lot         string
	}
	tests := []struct {
		name    string
		args    args
		want    *ModelNumber
		wantErr bool
	}{
		{"", args{"a20421", "", ""}, nil, true},
		{"", args{"", "100", ""}, nil, true},
		{"", args{"", "", "1"}, nil, true},
		{"", args{"a20421", "100", "1"}, &ModelNumber{"a20421", "100", "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := func() (got *ModelNumber, err error) {
				defer func() {
					e := recover()
					if e != nil {
						got = nil
						err = e.(error)
					}
				}()
				return NewModelNumber(tt.args.productCode, tt.args.branch, tt.args.lot), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewModelNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewModelNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelNumber_ModelNumber(t *testing.T) {
	type fields struct {
		productCode string
		branch      string
		lot         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"a20421", "100", "1"}, "a20421-100-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			modelNumber := &ModelNumber{
				productCode: tt.fields.productCode,
				branch:      tt.fields.branch,
				lot:         tt.fields.lot,
			}
			if got := modelNumber.ModelNumber(); got != tt.want {
				t.Errorf("ModelNumber.ModelNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
