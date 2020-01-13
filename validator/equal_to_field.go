package validator

import "fmt"

func (v *Validator) EqualToField(field string, value interface{}, toEqualField string, toEqualValue interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if value != toEqualValue {
		v.Errors[field] = fmt.Sprintf("%s must equal %s", field, toEqualField)
		return false
	}

	return true
}
