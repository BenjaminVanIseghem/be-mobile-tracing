package tracer

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

//Tracers is the slice that keeps all active tracers
var Tracers []opentracing.Tracer

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
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

// InitMultiple returns an array of instances of Jaeger Tracer that sample 100% of traces and log all spans to stdout.
func InitMultiple(services []string) ([]opentracing.Tracer, []io.Closer) {

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	for i := 0; i < len(services); i++ {
		tracer, closer, err := cfg.New(services[i], config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		}
		Tracers = append(Tracers, tracer)
		Closers = append(Closers, closer)
	}

	return Tracers, Closers
}

//AddNewTracer adds a new tracer to the existing Tracer and Closer slices
func AddNewTracer(service string) ([]opentracing.Tracer, []io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	Tracers = append(Tracers, tracer)
	Closers = append(Closers, closer)

	return Tracers, Closers
}

//CloseAllTracers closes all tracers in the given slice
func CloseAllTracers() {
	for i := 0; i < len(Closers); i++ {
		Closers[i].Close()
	}
}
