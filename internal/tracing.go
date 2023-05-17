package internal

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type TracingOpts struct {
	Name       string
	SpanName   string
	Attributes string
}

func Tracing(ctx context.Context, opts *TracingOpts) {
	tr := otel.Tracer(opts.Name)
	ctx, span := tr.Start(ctx, opts.SpanName)
	span.SetAttributes(attribute.Key("params").String(opts.Attributes))
	defer span.End()
}
