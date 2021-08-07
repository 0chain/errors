// Package errors - application error interface implementation
package errors

import (
	"fmt"
	"strings"
)

/*Error type for a new application error */
type Error struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg"`
}

/*
New - create a new error
two arguments can be passed!
1. code
2. message
if only one argument is passed its considered as message
if two arguments are passed then
	first argument is considered for code and
	second argument is considered for message
*/
func New(args ...string) *Error {
	return new(args...)
}

/*
Newf - creates a new error
*/
func Newf(code string, format string, args ...interface{}) *Error {
	return new(code, fmt.Sprintf(format, args...))
}

func (err *Error) Error() string {
	if err.Code == "" {
		return err.Msg
	}
	return fmt.Sprintf("%s: %s", err.Code, err.Msg)
}

func new(args ...string) *Error {
	currentError := Error{}

	switch len(args) {
	case 1:
		currentError.Msg = strings.TrimSpace(args[0])
	case 2:
		if isInvalidCode(args[0]) {
			return invalidCode(args[0])
		}
		currentError.Code, currentError.Msg = strings.TrimSpace(args[0]), strings.TrimSpace(args[1])
	default:
		return invalidUsage(args...)
	}

	return &currentError
}

func invalidUsage(args ...string) *Error {
	return &Error{
		Code: "incorrect_usage",
		Msg:  fmt.Sprintf("max allowed parameters is 2 i.e code, msg. parameters sent - %d", len(args)),
	}
}

func invalidCode(code string) *Error {
	return &Error{
		Code: "incorrect_code",
		Msg:  "code should not have spaces. use '" + strings.ToLower(strings.ReplaceAll(code, " ", "_")) + "' instead of '" + code + "'",
	}
}

func isInvalidCode(code string) bool {
	return len(strings.Split(code, " ")) != 1
}
