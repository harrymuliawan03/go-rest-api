package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = translate(v)
		}
	}

	return res
}

func translate(fe validator.FieldError) string {
	switch fe.ActualTag() {
	case "required":
		return fmt.Sprintf("Field %s is required", fe.StructField())
	case "email":
		return fmt.Sprintf("Field %s is not valid email", fe.StructField())
	case "unique":
		return fmt.Sprintf("Field %s must be unique", fe.StructField())
	case "min":
		return fmt.Sprintf("Field %s size minimum %s", fe.StructField(), fe.Param())
	}

	return "Validation failed"
}
