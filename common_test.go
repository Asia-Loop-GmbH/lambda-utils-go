package lambda_utils_go_test

import "testing"

func failedHandler(t *testing.T) func(message string, callerSkip ...int) {
	return func(message string, callerSkip ...int) {
		t.Error(message)
	}
}
