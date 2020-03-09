package value

import (
	"fmt"

	"github.com/trewanek/go-with-ddd-example/domain/derr"
)

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) *Money {
	if amount < 0.0 {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("amount should be positive value. value: %f", amount)))
	}

	if currency == "" {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("currency is required")))
	}

	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (money *Money) Add(arg *Money) (*Money, error) {
	if err := money.checkNil(arg); err != nil {
		return nil, err
	}
	if err := money.checkCurrency(arg); err != nil {
		return nil, err
	}

	return &Money{
		amount:   money.amount + arg.amount,
		currency: money.currency,
	}, nil
}

func (money *Money) Sub(arg *Money) (*Money, error) {
	if err := money.checkNil(arg); err != nil {
		return nil, err
	}
	if err := money.checkCurrency(arg); err != nil {
		return nil, err
	}
	if money.amount-arg.amount < 0 {
		return nil, fmt.Errorf(
			"calc result is negative. result should be positive. origin: %f, arg: %f",
			money.amount,
			arg.amount,
		)
	}

	return &Money{
		amount:   money.amount - arg.amount,
		currency: money.currency,
	}, nil
}

func (money *Money) checkCurrency(arg *Money) error {
	if money.currency != arg.currency {
		return derr.NewInValidArgumentErr(
			fmt.Errorf(
				"arg currency is not equal. orgin: %s, arg: %s",
				money.currency,
				arg.currency,
			),
		)
	}
	return nil
}

func (money *Money) checkNil(arg *Money) error {
	if arg == nil {
		return derr.NewInValidArgumentErr(
			fmt.Errorf("arg is required. arg is nil"),
		)
	}
	return nil
}
