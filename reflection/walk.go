// golang challenge: write a function walk(x interface{}, fn func(string))
// which takes a struct x and calls fn for all strings fields found inside.
// difficulty level: recursively.

package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// if it's a pointer, just work on the thing that it points to
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
