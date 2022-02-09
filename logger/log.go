package logger

import (
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
	"os"
)

type LogFields struct {
	Stage      *string
	Path       *string
	RequestID  *string
	HTTPMethod *string
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}

func NewLogger(f *LogFields) *log.Entry {
	return log.WithFields(log.Fields{
		"stage":      f.Stage,
		"path":       f.Path,
		"requestId":  f.RequestID,
		"httpMethod": f.HTTPMethod,
	})
}

func NewLoggerFromProxyRequest(f LogFields, request *events.APIGatewayProxyRequest) *log.Entry {
	return NewLogger(&LogFields{
		Stage:      &request.RequestContext.Stage,
		Path:       &request.Path,
		RequestID:  &request.RequestContext.RequestID,
		HTTPMethod: &request.HTTPMethod,
	})
}
