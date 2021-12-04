// golang challenge: write a function walk(x interface{}, fn func(string))
// which takes a struct x and calls fn for all strings fields found inside.
// difficulty level: recursively.

package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
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
