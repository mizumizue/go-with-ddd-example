package value_object

import (
	"reflect"
	"testing"
)

func TestNewMoney(t *testing.T) {
	type args struct {
		amount   float64
		currency string
	}
	tests := []struct {
		name    string
		args    args
		want    *Money
		wantErr bool
	}{
		{"", args{amount: 0.0, currency: "JPY"}, &Money{amount: 0.0, currency: "JPY"}, false},
		{"", args{amount: -1.0, currency: "JPY"}, nil, true},
		{"", args{amount: -1.0, currency: ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := func() (got *Money, err error) {
				defer func() {
					e := recover()
					if e != nil {
						got = nil
						err = e.(error)
					}
				}()
				return NewMoney(tt.args.amount, tt.args.currency), nil
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewModelNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMoney() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Add(t *testing.T) {
	type fields struct {
		amount   float64
		currency string
	}
	type args struct {
		arg *Money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Money
		wantErr bool
	}{
		{
			"",
			fields{100.0, "JPY"},
			args{&Money{200.0, "JPY"}},
			&Money{300.0, "JPY"},
			false,
		},
		{
			"",
			fields{100.0, "JPY"},
			args{&Money{100.0, "USD"}},
			nil,
			true,
		},
		{
			"",
			fields{100.0, "JPY"},
			args{nil},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := &Money{
				amount:   tt.fields.amount,
				currency: tt.fields.currency,
			}
			got, err := money.Add(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Money.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Sub(t *testing.T) {
	type fields struct {
		amount   float64
		currency string
	}
	type args struct {
		arg *Money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Money
		wantErr bool
	}{
		{
			"",
			fields{300.0, "JPY"},
			args{&Money{100.0, "JPY"}},
			&Money{200.0, "JPY"},
			false,
		},
		{
			"",
			fields{100.0, "JPY"},
			args{&Money{300.0, "USD"}},
			nil,
			true,
		},
		{
			"",
			fields{100.0, "JPY"},
			args{&Money{300.0, "JPY"}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := &Money{
				amount:   tt.fields.amount,
				currency: tt.fields.currency,
			}
			got, err := money.Sub(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Money.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_checkCurrency(t *testing.T) {
	type fields struct {
		amount   float64
		currency string
	}
	type args struct {
		arg *Money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"",
			fields{100.0, "JPY"},
			args{&Money{300.0, "JPY"}},
			false,
		},
		{
			"",
			fields{100.0, "JPY"},
			args{&Money{300.0, "USD"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := &Money{
				amount:   tt.fields.amount,
				currency: tt.fields.currency,
			}
			if err := money.checkCurrency(tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Money.checkCurrency() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
