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
	event := span.BaggageItem("eventId")
	file := span.BaggageItem("file")
	if event != "" && file != "" {
		logrus.WithField("eventId", event).WithField("file", file).Debug(s)
	} else if event != "" && file == "" {
		logrus.WithField("eventId", event).Debug(s)
	} else if event == "" && file != "" {
		logrus.WithField("file", file).Debug(s)
	} else {
		logrus.Debug(s)
	}

	return span
}

//Info adds info logs to the span and returns it
func Info(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Info", s),
	)

	//Check for baggage
	event := span.BaggageItem("eventId")
	file := span.BaggageItem("file")
	if event != "" && file != "" {
		logrus.WithField("eventId", event).WithField("file", file).Info(s)
	} else if event != "" && file == "" {
		logrus.WithField("eventId", event).Info(s)
	} else if event == "" && file != "" {
		logrus.WithField("file", file).Info(s)
	} else {
		logrus.Info(s)
	}

	return span
}

//Warning adds error logs to the span and returns it, also logs the error
func Warning(span opentracing.Span, s string) opentracing.Span {
	span.LogFields(
		log.String("Warning message", s),
	)

	//Check for baggage
	event := span.BaggageItem("eventId")
	file := span.BaggageItem("file")
	if event != "" && file != "" {
		logrus.WithField("eventId", event).WithField("file", file).Warning(s)
	} else if event != "" && file == "" {
		logrus.WithField("eventId", event).Warning(s)
	} else if event == "" && file != "" {
		logrus.WithField("file", file).Warning(s)
	} else {
		logrus.Warning(s)
	}

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
			event := span.BaggageItem("eventId")
			file := span.BaggageItem("file")
			if event != "" && file != "" {
				logrus.WithField("eventId", event).WithField("file", file).Errorf("%s: %v", s, err)
			} else if event != "" && file == "" {
				logrus.WithField("eventId", event).Errorf("%s: %v", s, err)
			} else if event == "" && file != "" {
				logrus.WithField("file", file).Errorf("%s: %v", s, err)
			} else {
				logrus.Errorf("%s: %v", s, err)
			}
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
	event := span.BaggageItem("eventId")
	file := span.BaggageItem("file")
	if event != "" && file != "" {
		logrus.WithField("eventId", event).WithField("file", file).Fatalf("%s: %v", s, err)
	} else if event != "" && file == "" {
		logrus.WithField("eventId", event).Fatalf("%s: %v", s, err)
	} else if event == "" && file != "" {
		logrus.WithField("file", file).Fatalf("%s: %v", s, err)
	} else {
		logrus.Fatalf("%s: %v", s, err)
	}

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
		event := span.BaggageItem("eventId")
		file := span.BaggageItem("file")
		if event != "" && file != "" {
			logrus.WithField("eventId", event).WithField("file", file).Errorf("%s: %v", s, i)
		} else if event != "" && file == "" {
			logrus.WithField("eventId", event).Errorf("%s: %v", s, i)
		} else if event == "" && file != "" {
			logrus.WithField("file", file).Errorf("%s: %v", s, i)
		} else {
			logrus.Errorf("%s: %v", s, i)
		}
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
		event := span.BaggageItem("eventId")
		file := span.BaggageItem("file")
		if event != "" && file != "" {
			logrus.WithField("eventId", event).WithField("file", file).Println(s, i)
		} else if event != "" && file == "" {
			logrus.WithField("eventId", event).Println(s, i)
		} else if event == "" && file != "" {
			logrus.WithField("file", file).Println(s, i)
		} else {
			logrus.Println(s, i)
		}
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
		event := span.BaggageItem("eventId")
		file := span.BaggageItem("file")
		if event != "" && file != "" {
			logrus.WithField("eventId", event).WithField("file", file).Println(s, obj)
		} else if event != "" && file == "" {
			logrus.WithField("eventId", event).Println(s, obj)
		} else if event == "" && file != "" {
			logrus.WithField("file", file).Println(s, obj)
		} else {
			logrus.Println(s, obj)
		}
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
