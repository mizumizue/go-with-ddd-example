package value_object

import (
	"fmt"
	"github.com/trewanek/go-with-ddd-example/domain/derr"
)

type ModelNumber struct {
	productCode string
	branch      string
	lot         string
}

func NewModelNumber(productCode string, branch string, lot string) (*ModelNumber, error) {
	if productCode == "" {
		return nil, derr.NewInValidArgumentErr(fmt.Errorf("productCode is required"))
	}
	if branch == "" {
		return nil, derr.NewInValidArgumentErr(fmt.Errorf("productCode is required"))
	}
	if lot == "" {
		return nil, derr.NewInValidArgumentErr(fmt.Errorf("productCode is required"))
	}
	return &ModelNumber{productCode: productCode, branch: branch, lot: lot}, nil
}

func (modelNumber *ModelNumber) ModelNumber() string {
	return fmt.Sprintf("%s-%s-%s", modelNumber.productCode, modelNumber.branch, modelNumber.lot)
}

func (modelNumber *ModelNumber) checkNil(arg *ModelNumber) error {
	if arg == nil {
		return derr.NewInValidArgumentErr(
			fmt.Errorf("arg is required. arg is nil"),
		)
	}
	return nil
}
