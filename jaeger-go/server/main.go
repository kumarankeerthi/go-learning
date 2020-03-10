package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kumarankeerthi/go-learning/jaeger-go/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

type Server struct {
	tracer opentracing.Tracer
}

func main() {
	fmt.Println("Jaeger sample server")
	tracer, closer := tracer.InitJaeger("Jaeger Server")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	s := &Server{
		tracer: tracer,
	}
	http.HandleFunc("/ping", s.handleRequest)

	http.ListenAndServe(":8060", nil)
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside HanlderFunc")
	fmt.Println(r.Method)
	spanCtc, _ := s.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := s.tracer.StartSpan("Server", ext.RPCServerOption(spanCtc))
	defer span.Finish()
	span.LogFields(
		log.String("jaeger-server", "serving /ping url"),
	)

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	// localspan, ctx := opentracing.StartSpanFromContext(ctx, "HanlderFunc")
	// defer localspan.Finish()
	fmt.Println("before")
	func2(ctx)
	func3(ctx)
	fmt.Println("after")
	w.Write([]byte("Working"))

}

func func2(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "func in server") //rootSpan.Tracer().StartSpan("func2", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	span.LogFields(
		log.String("event", "func in server"),
	)
	time.Sleep(2 * time.Second)
}

func func3(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "func3 in server") //rootSpan.Tracer().StartSpan("func2", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	span.LogFields(
		log.String("event", "func in server"),
	)
	time.Sleep(2 * time.Second)
}
