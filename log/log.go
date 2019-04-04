package log

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
)

//Debug adds debug logs to the span and returns it
func Debug(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Debug", s),
	)

	//Check for baggage
	fields := logrus.Fields{}
	span.Context().ForeachBaggageItem(func(key string, value string) bool {
		fields[key] = value
		return true
	})
	logrus.WithFields(fields).Debug(s)

	return span
}

//Info adds info logs to the span and returns it
func Info(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Info", s),
	)

	//Check for baggage
	fields := logrus.Fields{}
	span.Context().ForeachBaggageItem(func(key string, value string) bool {
		fields[key] = value
		return true
	})
	logrus.WithFields(fields).Info(s)

	return span
}

//Warning adds error logs to the span and returns it, also logs the error
func Warning(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Warning message", s),
	)

	//Check for baggage
	fields := logrus.Fields{}
	span.Context().ForeachBaggageItem(func(key string, value string) bool {
		fields[key] = value
		return true
	})
	logrus.WithFields(fields).Warning(s)

	return span
}

//Error adds error logs to the span and returns it, also logs the error
func Error(span opentracing.Span, err error, s string, isLog bool) opentracing.Span {
	span.LogFields(
		log.String("Error message", s),
		log.Error(err),
	)

	//Check for baggage and if it needs to be logged
	if err != nil {
		//Set error tag for filtering
		span.SetTag("error", true)
		if isLog {
			fields := logrus.Fields{}
			span.Context().ForeachBaggageItem(func(key string, value string) bool {
				fields[key] = value
				return true
			})
			logrus.WithFields(fields).Errorf("%s: %v", s, err)
		}
	}

	return span
}

//Fatal adds error logs to the span and returns it, also logs the error
func Fatal(span opentracing.Span, err error, s string) opentracing.Span {
	span.LogFields(
		log.String("Fatal message", s),
		log.Error(err),
	)

	//Set tags for filtering
	span.SetTag("error", true)
	span.SetTag("fatal", true)

	//Check for baggage
	fields := logrus.Fields{}
	span.Context().ForeachBaggageItem(func(key string, value string) bool {
		fields[key] = value
		return true
	})
	logrus.WithFields(fields).Fatalf("%s: %v", s, err)

	return span
}

//StatusCode adds statuscode logs to the span and returns it
func StatusCode(span opentracing.Span, s string, i int, isLog bool) opentracing.Span {
	span.LogFields(
		log.String("Message", s),
		log.Int("Statuscode", i),
	)
	//Set tag for filtering
	span.SetTag("error", true)

	//Check for baggage and if it needs to be logged
	if isLog {
		fields := logrus.Fields{}
		span.Context().ForeachBaggageItem(func(key string, value string) bool {
			fields[key] = value
			return true
		})
		logrus.WithFields(fields).Errorf("%s: %v", s, i)
	}

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
		fields := logrus.Fields{}
		span.Context().ForeachBaggageItem(func(key string, value string) bool {
			fields[key] = value
			return true
		})
		logrus.WithFields(fields).Println(s, i)
	}
	return span
}

//Object adds an object log to the span and returns it
func Object(span opentracing.Span, s string, obj interface{}, isLog bool) opentracing.Span {
	span.LogFields(
		log.Object(s, obj),
	)
	//Check if it needs to be logged
	if isLog {
		//Check for baggage
		fields := logrus.Fields{}
		span.Context().ForeachBaggageItem(func(key string, value string) bool {
			fields[key] = value
			return true
		})
		logrus.WithFields(fields).Errorf(s, obj)
	}
	return span
}

//StringMap maps the keys and values from a given map to span logs and returns the span
func StringMap(span opentracing.Span, m map[string]string) opentracing.Span {
	for key, value := range m {
		span.LogKV(key, value)
	}
	return span
}

//IntMap maps the keys and values from a given map to span logs and returns the span
func IntMap(span opentracing.Span, m map[string]int) opentracing.Span {
	for key, value := range m {
		span.LogKV(key, value)
	}
	return span
}

//InterfaceMap maps the keys and values from a given map to span logs and returns the span
func InterfaceMap(span opentracing.Span, m map[string]interface{}) opentracing.Span {
	for key, value := range m {
		span.LogKV(key, value)
	}
	return span
}
