package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type topTestCase struct {
	about              string
	err                error
	expectedTopMessage string
}

func getTopTestCases() []topTestCase {
	return []topTestCase{
		{
			about:              "error from errors package",
			err:                errors.New("error from errors package"),
			expectedTopMessage: "error from errors package",
		},
		{
			about:              "error from fmt package",
			err:                fmt.Errorf("error from fmt package"),
			expectedTopMessage: "error from fmt package",
		},
		{
			about:              "error from 0chain errors package",
			err:                New("", "error from 0chain errors package"),
			expectedTopMessage: "error from 0chain errors package",
		},
		{
			about:              "nil error",
			err:                nil,
			expectedTopMessage: "",
		},
	}
}

func TestTop(t *testing.T) {
	for _, gwtc := range getWrapTopTestCases() {
		t.Run(gwtc.about, func(t *testing.T) {
			var wrappedError error
			for _, tc := range gwtc.testCase {
				wrappedError = Wrap(wrappedError, tc)
			}
			require.Equal(t, gwtc.expectedTopMessage, Top(wrappedError))
		})
	}

	for _, gtc := range getTopTestCases() {
		t.Run(gtc.about, func(t *testing.T) {
			require.Equal(t, gtc.expectedTopMessage, Top(gtc.err))
		})
	}
}
