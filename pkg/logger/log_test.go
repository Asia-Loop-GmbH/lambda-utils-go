package logger_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
)

func TestNewEmptyLogger_Timestamp(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	l := logger.NewEmptyLogger()
	l.Info("this is a log message")
}
