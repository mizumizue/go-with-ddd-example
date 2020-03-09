package value

import (
	"fmt"

	"github.com/trewanek/go-with-ddd-example/domain/derr"
)

type ModelNumber struct {
	productCode string
	branch      string
	lot         string
}

func NewModelNumber(productCode string, branch string, lot string) *ModelNumber {
	if productCode == "" {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("productCode is required")))
	}
	if branch == "" {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("productCode is required")))
	}
	if lot == "" {
		panic(derr.NewInValidArgumentErr(fmt.Errorf("productCode is required")))
	}
	return &ModelNumber{productCode: productCode, branch: branch, lot: lot}
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
