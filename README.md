# Zipkin tracer DSL for Golang



## How to use:

Declare the environment variable ZIPKIN_URL with the protocol://host:port of the Zipkin collector.

Example: `http://127.0.0.1:9411`

Otherwise, it will use `http://127.0.0.1:9411` as default value.



## How to trace:

#### Trace your endpoints

```
func Handlers() http.Handler {
	r := mux.NewRouter()
	tracing.StartTracing("localhost:8080", "users")

	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler()).Name("/healthcheck")
	return r
}
err := http.ListenAndServe(":8080", tracing.TrackerHandler(Handlers()))
```

#### Trace your methods

```
trace := tracing.Trace("My API")
[YOUR CODE]
defer trace.Finish()
```

#### Trace your methods with parent

```
parent := tracing.Trace("Other API")
[YOUR CODE]
	span := tracing.TraceParent("Get Another API", parent)
	defer span.Finish()
defer parent.Finish()
```

