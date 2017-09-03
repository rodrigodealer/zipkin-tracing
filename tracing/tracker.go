package tracing

import (
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func TrackerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wireContext, _ := opentracing.GlobalTracer().Extract(
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		span := opentracing.StartSpan(r.RequestURI, ext.RPCServerOption(wireContext))
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(r.Context(), span)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
