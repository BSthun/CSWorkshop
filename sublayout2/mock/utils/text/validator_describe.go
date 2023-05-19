package text

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func DescribeValidator(err validator.FieldError) string {
	var typ string
	if strings.Contains(err.Kind().String(), "int") {
		typ = "%d"
	}
	if strings.Contains(err.Kind().String(), "float") {
		typ = "%.2f"
	}
	if err.Kind() == reflect.String {
		typ = "%s"
	}

	if err.Tag() == "required" {
		return fmt.Sprintf("%s is required", err.Field())
	}
	if err.Tag() == "email" {
		return fmt.Sprintf("%s (%f) is not a valid email", err.Field(), err.Value())
	}
	if err.Tag() == "lte" {
		return fmt.Sprintf("%s ("+typ+") must be less than or equal to %s", err.Field(), err.Value(), err.Param())
	}
	if err.Tag() == "gte" {
		return fmt.Sprintf("%s ("+typ+") must be greater than or equal to %s", err.Field(), err.Value(), err.Param())
	}
	if err.Tag() == "lt" {
		return fmt.Sprintf("%s ("+typ+") must be less than %s", err.Field(), err.Value(), err.Param())
	}
	if err.Tag() == "gt" {
		return fmt.Sprintf("%s ("+typ+") must be greater than %s", err.Field(), err.Value(), err.Param())
	}
	if err.Tag() == "eq" {
		return fmt.Sprintf("%s ("+typ+") must be equal to %s", err.Field(), err.Value(), err.Param())
	}
	if err.Tag() == "ne" {
		return fmt.Sprintf("%s ("+typ+") must not be equal to %s", err.Field(), err.Value(), err.Param())
	}
	if err.Tag() == "url" {
		return fmt.Sprintf("%s (%s) must be a valid URL", err.Field(), err.Value())
	}
	return fmt.Sprintf("%s (%s) is invalid %s %s", err.Field(), err.Value(), err.Tag(), err.Param())
}
