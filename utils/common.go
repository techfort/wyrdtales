package utils

import (
	"fmt"
)

// PanicIf panics if error is not null
func PanicIf(err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("Panicking with error: %+v", err.Error()))
		panic(err)
	}
}

// Bypass panics if there's an error and returns a good result
func Bypass(i interface{}, err error) interface{} {
	PanicIf(err)
	return i
}
