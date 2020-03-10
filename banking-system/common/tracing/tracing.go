package tracing

import (
	"context"
	"fmt"
	"io"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

var tracer opentracing.Tracer
var closer io.Closer

//InitializeTracing will initiazie tracing for a give service name
func InitializeTracing(serviceName string) {
	cfg := &jaegerconfig.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans: true,
		},
	}
	trc, cls, err := cfg.NewTracer(jaegerconfig.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Println("Error initialiazing Tracing for servicename :", serviceName)
	}
	tracer = trc
	closer = cls
	opentracing.SetGlobalTracer(tracer)
	//	return tracer, closer
}

// Trace is the middleware that for a given handler it will publish spans
func Trace(methodName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			carrier := opentracing.HTTPHeadersCarrier(r.Header)
			spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, carrier)
			var span opentracing.Span
			if err == nil {
				span = tracer.StartSpan(methodName, ext.RPCServerOption(spanCtx))
			} else {
				span = tracer.StartSpan(methodName)
			}

			defer span.Finish()
			ctx := opentracing.ContextWithSpan(r.Context(), span)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func TraceFuncCall(funcName string, ctx context.Context) opentracing.Span {
	span, _ := opentracing.StartSpanFromContext(ctx, funcName)
	return span
}
