package value

import "reflect"

func StructSize(x any) int {
	v := reflect.ValueOf(x)
	return v.NumField()
}

func StructIndex[T any](x any, index int) T {
	v := reflect.ValueOf(x)
	return v.Field(index).Interface().(T)
}
