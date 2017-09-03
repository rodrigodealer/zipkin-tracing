package tracing

import opentracing "github.com/opentracing/opentracing-go"

func Trace(title string) *Tracing {
	tracing := &Tracing{Span: opentracing.StartSpan(title)}
	return tracing
}

func TraceParent(title string, parent *Tracing) *Tracing {
	if parent.Span != nil {
		var tracing = &Tracing{Span: opentracing.StartSpan(title, opentracing.ChildOf(parent.Span.Context()))}
		return tracing
	}
	return &Tracing{}
}

func (t *Tracing) Finish() {
	if t.Span != nil {
		t.Span.Finish()
	}
}

type Tracing struct {
	Span opentracing.Span
}
