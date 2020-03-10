package tracer

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegerconfig.Configuration{
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, jaegerconfig.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Println("error creating tracer")
	}
	return tracer, closer
}
