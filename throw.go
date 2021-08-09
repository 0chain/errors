package errors

import "strings"

// ApplicationError an appliction error with predifined error variable and detail message
type ApplicationError struct {
	// Inner inner error
	Inner error
	// MsgList detail message
	MsgList []string
}

// Error implement error.Error
func (e *ApplicationError) Error() string {
	if e == nil {
		return ""
	}

	if e.Inner != nil && e.MsgList != nil {
		return e.Inner.Error() + ": " + strings.Join(e.MsgList, " ")
	} else if e.Inner != nil {
		return e.Inner.Error()
	} else {
		return strings.Join(e.MsgList, " ")
	}

}

// Unwrap implement error.Unwrap
func (e *ApplicationError) Unwrap() error {
	if e == nil {
		return nil
	}

	if e.Inner != nil {
		return e.Inner
	}

	return e
}

// Throw create an application error with prefinded error variable and message
// example
//    errors.Throw(ErrInvalidParameter, "bloober_id is missing")
func Throw(inner error, msgList ...string) error {
	return &ApplicationError{
		Inner:   inner,
		MsgList: msgList,
	}
}
