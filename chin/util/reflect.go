package util

import (
	"reflect"
)

// GetPointerTo takes an interface that may be a pointer to a value and returns an
// interface that's a pointer to the value.
func GetPointerTo(i any) any {
	v := reflect.ValueOf(i)

	// check if i is already a pointer
	if v.Kind() == reflect.Ptr {
		return i
	}

	// create a new pointer to v
	p := reflect.New(v.Type())
	p.Elem().Set(v)
	return p.Interface()
}

// GetPointedTo is a function that an interface that may be a pointer to a value and
// returns an interface that's the value.
func GetPointedTo(i any) any {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr {
		return i
	}

	return v.Elem().Interface()
}
