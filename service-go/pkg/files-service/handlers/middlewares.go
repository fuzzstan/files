package handlers

import (
	"github.com/chaos-io/gokit/middleware"
	"github.com/chaos-io/gokit/tracing"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"go.opentelemetry.io/otel/trace"

	"github.com/fuzzstan/files/go/pkg/files"
	"github.com/fuzzstan/files/service-go/pkg/files-service/svc"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	// this service api
	pb "github.com/fuzzstan/files/go/pkg/files/v1"
)

var (
	_ = core.Timestamp{}
	_ = files.File{}
	_ = core.Null{}
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/ncraft-io//_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	var tracer trace.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(trace.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	// var validator *middleware.Validator
	// if value, ok := options["validator"]; ok && value != nil {
	// 	validator = value.(*middleware.Validator)
	// }

	{ // get_files
		if tracer != nil {
			in.GetFilesEndpoint = tracing.TraceServer(tracer, "get_files")(in.GetFilesEndpoint)
		}
		if count != nil && latency != nil {
			in.GetFilesEndpoint = middleware.Instrumenting(latency.With("method", "get_files"), count.With("method", "get_files"))(in.GetFilesEndpoint)
		}
		// if validator != nil {
		// 	in.GetFilesEndpoint = validator.Validate()(in.GetFilesEndpoint)
		// }
	}
	{ // get_file
		if tracer != nil {
			in.GetFileEndpoint = tracing.TraceServer(tracer, "get_file")(in.GetFileEndpoint)
		}
		if count != nil && latency != nil {
			in.GetFileEndpoint = middleware.Instrumenting(latency.With("method", "get_file"), count.With("method", "get_file"))(in.GetFileEndpoint)
		}
		// if validator != nil {
		// 	in.GetFileEndpoint = validator.Validate()(in.GetFileEndpoint)
		// }
	}
	{ // create_file
		if tracer != nil {
			in.CreateFileEndpoint = tracing.TraceServer(tracer, "create_file")(in.CreateFileEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateFileEndpoint = middleware.Instrumenting(latency.With("method", "create_file"), count.With("method", "create_file"))(in.CreateFileEndpoint)
		}
		// if validator != nil {
		// 	in.CreateFileEndpoint = validator.Validate()(in.CreateFileEndpoint)
		// }
	}
	{ // update_file
		if tracer != nil {
			in.UpdateFileEndpoint = tracing.TraceServer(tracer, "update_file")(in.UpdateFileEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateFileEndpoint = middleware.Instrumenting(latency.With("method", "update_file"), count.With("method", "update_file"))(in.UpdateFileEndpoint)
		}
		// if validator != nil {
		// 	in.UpdateFileEndpoint = validator.Validate()(in.UpdateFileEndpoint)
		// }
	}
	{ // delete_file
		if tracer != nil {
			in.DeleteFileEndpoint = tracing.TraceServer(tracer, "delete_file")(in.DeleteFileEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteFileEndpoint = middleware.Instrumenting(latency.With("method", "delete_file"), count.With("method", "delete_file"))(in.DeleteFileEndpoint)
		}
		// if validator != nil {
		// 	in.DeleteFileEndpoint = validator.Validate()(in.DeleteFileEndpoint)
		// }
	}

	return in
}

func WrapService(in pb.FilesServer, options map[string]interface{}) pb.FilesServer {
	return in
}
