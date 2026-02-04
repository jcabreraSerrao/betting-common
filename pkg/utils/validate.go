package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

type CustomValidator struct {
	validator.Validate
}

func (v *CustomValidator) DecimalGt(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}
	min, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return data.GreaterThan(min)
}

func (v *CustomValidator) DecimalMin(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}
	min, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return data.GreaterThanOrEqual(min)
}

func NewCustomValidator() *CustomValidator {
	validate := &CustomValidator{
		Validate: *validator.New(),
	}
	_ = validate.RegisterValidation("d_min", validate.DecimalMin)
	_ = validate.RegisterValidation("d_gt", validate.DecimalGt)

	return validate
}
