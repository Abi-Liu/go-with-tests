package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		kind := field.Kind()
		if kind == reflect.String {
			fn(field.String())
		}

		if kind == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
