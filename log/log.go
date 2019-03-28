package log

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

//Error adds error logs to the spans and returns it
func Error(span opentracing.Span, err error, s string) opentracing.Span {
	span.LogFields(
		log.String("Error message", s),
		log.Error(err),
	)
	return span
}

//StatusCode adds statuscode logs to the spans and returns it
func StatusCode(i int) {

}

//String adds string logs to the spans and returns it
func String(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Message", s),
	)
	return span
}

//Int adds integer logs to the spans and returns it
func Int(i int) {

}
