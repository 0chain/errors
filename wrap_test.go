package errors

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type wrapTopTestCase struct {
	about              string
	testCase           []interface{}
	expectedTopMessage string
}

var nilWrapError error = nil

func getWrapTopTestCases() []wrapTopTestCase {
	return []wrapTopTestCase{
		{
			about: "wrapping all errors",
			testCase: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				New("", "This is a very big error! Beware of it!"),
				New("401", ""),
				errors.New("error created from err package"),
				fmt.Errorf("%s", "error created from fmt package"),
				nil,
			},
			expectedTopMessage: "incorrect_usage: you should pass either error or message to properly wrap the error! - wrapped with unsupported type, it should be either of type error/string",
		},
		{
			about: "wrapping all messages",
			testCase: []interface{}{
				"This is a very \"big\" error! Beware of it!",
				"This is a very 'big' error! Beware of it!",
				"This is a short error!",
				"",
			},
			expectedTopMessage: "incorrect_usage: you should pass either error or message to properly wrap the error! - wrapped with empty string",
		},
		{
			about: "wrapping errors and messages",
			testCase: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				"This is a very \"big\" error! Beware of it!",
				New("401", ""),
				"This is a very 'big' error! Beware of it!",
				"",
				nil,
				"This is a short error!",
				New("", "This is a very big error! Beware of it!"),
			},
			expectedTopMessage: "This is a very big error! Beware of it!",
		},
		{
			about: "wrapping error with nil error",
			testCase: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				nil,
				errors.New(""),
			},
			expectedTopMessage: "",
		},
		{
			about: "wrapping error with nil error",
			testCase: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				nil,
				&nilWrapError,
			},
			expectedTopMessage: "incorrect_usage: you should pass either error or message to properly wrap the error! - wrapped with unsupported type, it should be either of type error/string",
		},
	}
}

func TestWrap(t *testing.T) {
	for _, gtc := range getWrapTopTestCases() {
		t.Run(gtc.about, func(t *testing.T) {
			var wrappedError error
			for _, tc := range gtc.testCase {
				wrappedError = Wrap(wrappedError, tc)
			}

			require.Equal(t, len(gtc.testCase), len(strings.Split(wrappedError.Error(), "\n")))
		})
	}
}
