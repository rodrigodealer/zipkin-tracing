package tracing

import (
	"fmt"
	"log"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const zipkinHTTPEndpoint = "http://127.0.0.1:9411/api/v1/spans"

func zipkinEndpoint() string {
	var endpoint = zipkinHTTPEndpoint
	if os.Getenv("ZIPKIN_URL") == "" {
		endpoint = os.Getenv("ZIPKIN_URL")
	}
	return fmt.Sprintf("%v/api/v1/spans", endpoint)
}

func StartTracing(server string, service string) {
	collector, err := zipkin.NewHTTPCollector(zipkinEndpoint())
	if err != nil {
		log.Printf("unable to create Zipkin HTTP collector: %+v", err)
	}

	recorder := zipkin.NewRecorder(collector, false, server, service)

	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(true),
	)

	if err != nil {
		log.Printf("unable to create Zipkin tracer: %+v", err)
	}

	opentracing.InitGlobalTracer(tracer)
}
