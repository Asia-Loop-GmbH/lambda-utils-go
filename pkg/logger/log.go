package logger

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

type LogFields struct {
	Stage      *string
	Path       *string
	RequestID  *string
	HTTPMethod *string
}

func init() {
	var formatter log.Formatter
	jsonFormatter := &log.JSONFormatter{}
	jsonFormatter.TimestampFormat = "2006-01-02T15:04:05.999Z07:00"
	textFormatter := &log.TextFormatter{}
	local := os.Getenv("AL_LOCAL")
	if local == "true" {
		formatter = textFormatter
	} else {
		formatter = jsonFormatter
	}
	log.SetFormatter(formatter)
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

func NewEmptyLogger() *log.Entry {
	return log.WithFields(log.Fields{})
}

func NewLoggerFromProxyRequest(request *events.APIGatewayProxyRequest) *log.Entry {
	return NewLogger(&LogFields{
		Stage:      &request.RequestContext.Stage,
		Path:       &request.Path,
		RequestID:  &request.RequestContext.RequestID,
		HTTPMethod: &request.HTTPMethod,
	})
}
