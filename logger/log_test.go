package logger_test

import (
	"testing"

	. "github.com/onsi/gomega"

	logger2 "github.com/asia-loop-gmbh/lambda-utils-go/v2/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v2/test"
)

func TestNewEmptyLogger_Timestamp(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	logger := logger2.NewEmptyLogger()
	logger.Info("this is a log message")
}
