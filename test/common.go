package test

import "testing"

func FailedHandler(t *testing.T) func(message string, callerSkip ...int) {
	return func(message string, callerSkip ...int) {
		t.Error(message)
	}
}
