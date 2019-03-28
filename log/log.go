package log

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func logError(span opentracing.Span, err error, s string) opentracing.Span {
	span.LogFields(
		log.String("Error message", s),
		log.Error(err),
	)
	return span
}

func logStatusCode(i int) {

}

func logString(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Message", s),
	)
	return span
}

func logInt(i int) {

}
