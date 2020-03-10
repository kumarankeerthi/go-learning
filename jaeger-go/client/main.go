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

func main() {

	fmt.Println("Jaeger Go sample application")
	tracer, closer := tracer.InitJaeger("Jaeger Client")
	ctx := context.Background()
	defer closer.Close()
	//span := tracer.StartSpan("sample")
	opentracing.SetGlobalTracer(tracer)
	span, ctx := opentracing.StartSpanFromContext(ctx, "Main Function")

	span.SetTag("Client", "Main Function")

	defer span.Finish()

	span.LogFields(
		log.String("event", "Starting Main function"),
	)

	func1(ctx)

	httpClient := http.Client{}
	url := "http://localhost:8060/ping"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error creating request")
	}
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("error from server")
	}
	fmt.Println(resp.Body)
	func3(ctx)
	func4(ctx)
	http.Get(url)

}
func func3(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "func3") //rootSpan.Tracer().StartSpan("func2", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	span.LogFields(
		log.String("event", "func2"),
	)
	time.Sleep(3 * time.Second)
}
func func4(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "func4") //rootSpan.Tracer().StartSpan("func1", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	span.LogFields(
		log.String("event", "Strating func1"),
	)
	time.Sleep(4 * time.Second)

}

func func2(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "func2") //rootSpan.Tracer().StartSpan("func2", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	span.LogFields(
		log.String("event", "func2"),
	)
	time.Sleep(2 * time.Second)
}
func func1(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "func1") // rootSpan.Tracer().StartSpan("func1", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	span.LogFields(
		log.String("event", "Strating func1"),
	)
	time.Sleep(1 * time.Second)
	func2(ctx)

}
