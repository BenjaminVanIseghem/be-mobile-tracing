package log

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
)

//Future functionality will include DEBUG and WARNING messages which can be filtered or excluded in jaeger
//IDEA: ADD EXTRA PARAMETER TO PASS WHICH LEVEL SHOULD BE LOGGED (DEBUG, INFO, WARNING, ERROR, FATAL, PANIC, OR NONE)

//Error adds error logs to the span and returns it, also logs the error
func Error(span opentracing.Span, err error, s string, isLog bool) opentracing.Span {
	span.LogFields(
		log.String("Error message", s),
		log.Error(err),
	)

	span.SetTag("Error", err.Error())

	if isLog {
		logrus.Errorf("%s: %v", s, err)
	}

	return span
}

//Fatal adds error logs to the span and returns it, also logs the error
func Fatal(span opentracing.Span, err error, s string) opentracing.Span {
	span.LogFields(
		log.String("Fatal message", s),
		log.Error(err),
	)

	span.SetTag("Fatal", err.Error())

	logrus.Fatalf("%s: %v", s, err)

	return span
}

//StatusCode adds statuscode logs to the span and returns it
func StatusCode(span opentracing.Span, s string, i int, isLog bool) opentracing.Span {
	span.LogFields(
		log.String("Message", s),
		log.Int("Statuscode", i),
	)
	if isLog {
		logrus.WithField("Statuscode", i).Error(s, i)
	}

	return span
}

//Info adds info logs to the span and returns it
func Info(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Info", s),
	)
	logrus.Info(s)
	return span
}

//String adds string logs to the span and returns it
func String(span opentracing.Span, s string, s2 string) opentracing.Span {
	span.LogFields(
		log.String(s, s2),
	)

	return span
}

//Int adds integer logs to the span and returns it
func Int(span opentracing.Span, s string, i int, isLog bool) opentracing.Span {
	span.LogFields(
		log.Int(s, i),
	)
	if isLog {
		logrus.Println(s, i)
	}
	return span
}

//Object adds an object log to the span and returns it
func Object(span opentracing.Span, s string, obj interface{}, isLog bool) opentracing.Span {
	span.LogFields(
		log.Object(s, obj),
	)
	if isLog {
		logrus.Println(s, obj)
	}
	return span
}
