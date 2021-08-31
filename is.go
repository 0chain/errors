// Package errors - Is
package errors

import (
	"errors"
	"fmt"
	"reflect"
)

/*Is - tells whether actual error is targer error
where, actual error can be either Error/withError
if actual error is wrapped error then if any internal error
matches the target error then function results in true*/
func Is(actual error, target error) bool {

	fmt.Println("Is ..................................... 1")

	fmt.Println("actual error:", actual)
	fmt.Println("actual error kind:", reflect.ValueOf(actual).Kind())
	fmt.Println("target error:", target)
	fmt.Println("target error kind:", reflect.ValueOf(target).Kind())
	fmt.Println("Is ..................................... 2")

	if errors.Is(actual, target) {
		fmt.Println("Is ..................................... 3")
		return true
	}

	switch targetError := target.(type) {
	case *Error:
		fmt.Println("Is ..................................... 4")
		switch actualError := actual.(type) {
		case *Error:
			fmt.Println("Is ..................................... 5")
			if actualError.Code == "" && targetError.Code == "" {
				fmt.Println("Is ..................................... 5.1")
				return actualError.Msg == targetError.Msg
			}

			fmt.Println("Is ..................................... 5.2")
			fmt.Println("actualError.Code: ", actualError.Code)
			fmt.Println("actualError.Msg: ", actualError.Msg)
			fmt.Println("targetError.Code: ", targetError.Code)
			fmt.Println("targetError.Msg: ", targetError.Msg)

			fmt.Println("actualError.Code == targetError.Code: ", actualError.Code == targetError.Code)
			return actualError.Code == targetError.Code

		case *withError:
			fmt.Println("Is ..................................... 6")
			return Is(actualError.current, target) || Is(actualError.previous, target)
		default:
			fmt.Println("Is ..................................... 7")
			return false
		}
	default:
		fmt.Println("Is ..................................... 8")
		return false
	}
}

// As wrap errors.As
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Unwrap wrap errors.Unwrap
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
