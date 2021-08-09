package errors

// ApplicationError an appliction error with predifined error variable and detail message
type ApplicationError struct {
	// Inner inner error
	Inner error
	// Message detail message
	Message string
}

// Error implement error.Error
func (e *ApplicationError) Error() string {
	if e == nil {
		return ""
	}

	if e.Inner != nil && e.Message != "" {
		return e.Inner.Error() + ": " + e.Message
	} else if e.Inner != nil {
		return e.Inner.Error()
	} else {
		return e.Message
	}

}

// Unwrap implement error.Unwrap
func (e *ApplicationError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Inner
}

// Throw create an application error with prefinded error variable and message
// example
//    errors.Throw(ErrInvalidParameter, "bloober_id is missing")
func Throw(inner error, message string) error {
	return &ApplicationError{
		Inner:   inner,
		Message: message,
	}
}
