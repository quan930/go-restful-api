package util

import (
	"bytes"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateErrorFormat(err error) string {
	var buffer bytes.Buffer
	if _, ok := err.(*validator.InvalidValidationError); ok {
		buffer.WriteString(err.Error())
		return buffer.String()
	}

	for _, err := range err.(validator.ValidationErrors) {
		buffer.WriteString(err.Field())
		buffer.WriteString("参数异常,")
	}
	return strings.TrimRight(buffer.String(), ",")
}