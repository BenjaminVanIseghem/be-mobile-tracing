package log

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

//LogError adds error logs to the spans and returns it
func LogError(span opentracing.Span, err error, s string) opentracing.Span {
	span.LogFields(
		log.String("Error message", s),
		log.Error(err),
	)
	return span
}

//LogStatusCode adds statuscode logs to the spans and returns it
func LogStatusCode(i int) {

}

//LogString adds string logs to the spans and returns it
func LogString(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Message", s),
	)
	return span
}

//LogInt adds integer logs to the spans and returns it
func LogInt(i int) {

}
