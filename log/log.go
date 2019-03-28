package log

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

//Future functionality will include DEBUG and WARNING messages which can be filtered or excluded in jaeger
//LOGRUS support will be implemented if necessary (to eliminate duplicate code)

//Error adds error logs to the span and returns it
func Error(span opentracing.Span, err error, s string) opentracing.Span {
	span.LogFields(
		log.String("Error message", s),
		log.Error(err),
	)
	return span
}

//StatusCode adds statuscode logs to the span and returns it
func StatusCode(span opentracing.Span, i int) opentracing.Span {
	span.LogFields(
		log.Int("Statuscode", i),
	)
	return span
}

//String adds string logs to the span and returns it
func String(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Info", s),
	)
	return span
}

//Int adds integer logs to the span and returns it
func Int(span opentracing.Span, s string, i int) opentracing.Span {
	span.LogFields(
		log.Int(s, i),
	)
	return span
}

//Object adds an object log to the span and returns it
func Object(span opentracing.Span, s string, obj interface{}) opentracing.Span {
	span.LogFields(
		log.Object(s, obj),
	)

	return span
}
