package otelx

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

func Resource(attrs ...attribute.KeyValue) *resource.Resource {
	attributes := resource.Default().Attributes()
	attributes = append(attributes, attrs...)
	return resource.NewSchemaless(attributes...)
}
