package value

import (
	"reflect"
)

func Map(s any, tag string) map[string]any {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	m := map[string]any{}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			t := v.Type().Field(i).Tag.Get(tag)
			if t == "-" {
				continue
			}
			m[t] = v.Field(i).Interface()
		}
	}

	return m
}

func MapKey[T comparable](m map[T]any, key T) bool {
	_, ok := m[key]
	return ok
}
