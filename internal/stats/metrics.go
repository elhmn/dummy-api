package stats

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func SetupMetrics(ctx context.Context, serviceName string) (*metric.MeterProvider, error) {
	exporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, err
	}

	// labels/tags/resources that are common to all metrics.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		attribute.String("test", "elhmn"),
	)

	mp := metric.NewMeterProvider(
		metric.WithResource(resource),
		metric.WithReader(
			metric.NewPeriodicReader(exporter, metric.WithInterval(1*time.Second)),
		),
	)

	global.SetMeterProvider(mp)

	return mp, nil
}

// This was implemented using the uptrace go tutorial on opentelemetry
// https://uptrace.dev/opentelemetry/go-metrics.html#introduction
// and https://uptrace.dev/opentelemetry/metrics.html

// Meter can be a global/package variable.
var Meter = global.MeterProvider().Meter("dummy-api")

// Metrics
const (
	HTTPServerRequests      = "http.server.requests"
	HTTPServerRequestsWrite = "http.server.requests.write"
	HTTPServerRequestsRead  = "http.server.requests.read"
	HTTPServerLatency       = "http.server.latency"
	HTTPServerErrorsWrite   = "http.server.errors.write"
	HTTPServerErrorsRead    = "http.server.errors.read"
	HTTPServerOKWrite       = "http.server.ok.write"
	HTTPServerOKRead        = "http.server.ok.read"
)

func IncrementHTTPServerRequest(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerRequests,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}

func IncrementHTTPServerWriteRequest(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerRequestsWrite,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http write requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}

func IncrementHTTPServerReadRequest(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerRequestsRead,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http read requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}

func IncrementHTTPServerErrorWrite(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerErrorsWrite,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http write requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}

func IncrementHTTPServerErrorRead(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerErrorsRead,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http read requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}

func IncrementHTTPServerOKWrite(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerOKWrite,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http write requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}

func IncrementHTTPServerOKRead(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		HTTPServerOKRead,
		instrument.WithUnit("1"),
		instrument.WithDescription("number of http read requests received"),
	)

	// Increment the counter.
	counter.Add(ctx, 1)
}
