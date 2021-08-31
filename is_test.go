package errors

import (
	"errors"
	"fmt"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	ErrInvalid = fs.ErrInvalid // "invalid argument"
	ErrInternal = New("internal_error", "this is a internal error")
)

type isErrorWrapTestCase struct {
	about        string
	preCondition []interface{}
	testCase     map[error]bool
}

func getIsErrorWrapTestCases() []isErrorWrapTestCase {
	return []isErrorWrapTestCase{
		{
			about: "wrapping all errors",
			preCondition: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				New("", "This is a very big error! Beware of it!"),
				New("401", ""),

				errors.New("error created from err package"),
				fmt.Errorf("%s", "error created from fmt package"),
				nil,
			},
			testCase: map[error]bool{
				New("500", "This is a very big error! Beware of it!"): true,
			},
		},
		{
			about: "wrapping error with nil error",
			preCondition: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				nil,
				errors.New(""),
			},
			testCase: map[error]bool{
				New("500", "This is a very big error! Beware of it!"): true,

				new("500", "This is a very big error! Beware of it!!!!!"): true,
			},
		},
		{
			about: "wrapping error with nil error",
			preCondition: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
			},
			testCase: map[error]bool{
				New("500", "This is a very big error! Beware of it!"): true,
			},
		},
	}
}

type isErrorTestCase struct {
	about       string
	actualError error
	targetError error
	expected    bool
}

func isErrorTestCases() []isErrorTestCase {
	return []isErrorTestCase{
		{
			about:       "actual error and target Error",
			actualError: errors.New("actual error"),
			targetError: New("", "actual error"),
			expected:    false,
		},
		{
			about:       "actual error and target error",
			actualError: ErrInvalid,
			targetError: ErrInvalid,
			expected:    true,
		},
		{
			about:       "actual error and target nil",
			actualError: ErrInvalid,
			targetError: nil,
			expected:    false,
		},
		{
			about:       "actual nil and target error",
			actualError: nil,
			targetError: ErrInvalid,
			expected:    false,
		},
		{
			about:       "actual withError and target error",
			actualError: Wrap(Wrap(New("", "actual error"), "wrapped error1"), "wrapped error2"),
			targetError: New("", "wrapped error1"),
			expected:    true,
		},
		{
			about: "internal error",
			actualError: Wrap(ErrInternal, "internal server error"),
			targetError: ErrInternal,
			expected: true,
		},
	}
}

func TestIs(t *testing.T) {
	for _, gtc := range getIsErrorWrapTestCases() {
		t.Run(gtc.about, func(t *testing.T) {
			var wrappedError error
			for _, tc := range gtc.preCondition {
				wrappedError = Wrap(wrappedError, tc)
			}

			for k, v := range gtc.testCase {
				require.Equal(t, v, Is(wrappedError, k))
			}
		})
	}

	for _, iet := range isErrorTestCases() {
		t.Run(iet.about, func(t *testing.T) {
			require.Equal(t, iet.expected, Is(iet.actualError, iet.targetError))
		})
	}

}
