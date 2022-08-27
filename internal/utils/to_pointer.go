package utils

import "reflect"

type Constraint interface {
	bool | int | string | int16 | int32 | int64 | float32 | float64 | interface{}
}

func ToPointer[T Constraint](value T) *T {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		panic("you obviously don't mean to get a pointer to a pointer.")
	}

	return &value
}
