package tracer

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

//Tracers is the map that keeps all active tracers with their names as the key
var Tracers = map[string]opentracing.Tracer{}

//Closers is the slice that keeps the closers to close all tracers
var Closers []io.Closer

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: false,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

// InitMultiple returns an array of instances of Jaeger Tracer that sample 100% of traces and log all spans to stdout.
func InitMultiple(services []string) {

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: false,
		},
	}
	for i := 0; i < len(services); i++ {
		tracer, closer, err := cfg.New(services[i], config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		}
		Tracers[services[i]] = tracer
		Closers = append(Closers, closer)
	}
}

//AddNewTracer adds a new tracer to the existing Tracer and Closer slices
func AddNewTracer(service string) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: false,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	Tracers[service] = tracer
	Closers = append(Closers, closer)
}

//CloseAllTracers closes all tracers in the given slice
func CloseAllTracers() {
	for i := 0; i < len(Closers); i++ {
		Closers[i].Close()
	}
}
