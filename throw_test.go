package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThrowIs(t *testing.T) {

	errIs := errors.New("Is")

	errIsNot := errors.New("Is not")

	err := Throw(errIs, "", "errors.Is works")

	require.ErrorIs(t, err, errIs)
	require.Equal(t, false, errors.Is(err, errIsNot))
}

func TestThrowUnwrap(t *testing.T) {

	inner := errors.New("Inner")

	err := Throw(inner, "errors.Unwrap works")

	require.Equal(t, inner, errors.Unwrap(err))
}

func TestThrowAs(t *testing.T) {

	inner := errors.New("Inner")

	err := Throw(inner, "errors.As works")

	var appErr *ApplicationError

	require.Equal(t, true, errors.As(err, &appErr))
}
