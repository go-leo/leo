package resourcex

import (
	"context"

	"github.com/go-leo/gox/stringx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

type ResourceFlag int

const (
	ResEnv ResourceFlag = 1 << iota
	ResHost
	ResTelemetrySDK
	ResOS
	ResProcess
	ResContainer
)

type Service struct {
	Name       string
	Namespace  string
	InstanceID string
	Version    string
}

func (svc *Service) Attributes() []attribute.KeyValue {
	if svc == nil {
		return nil
	}
	var attrs []attribute.KeyValue
	if stringx.IsNotBlank(svc.Name) {
		attrs = append(attrs, attribute.Key("service.name").String(svc.Name))
	}
	if stringx.IsNotBlank(svc.Namespace) {
		attrs = append(attrs, attribute.Key("service.namespace").String(svc.Namespace))
	}
	if stringx.IsNotBlank(svc.InstanceID) {
		attrs = append(attrs, attribute.Key("service.instance.id").String(svc.InstanceID))
	}
	if stringx.IsNotBlank(svc.Version) {
		attrs = append(attrs, attribute.Key("service.version").String(svc.Version))
	}
	return attrs
}

func NewResource(ctx context.Context, svc *Service, res ResourceFlag, attrs ...attribute.KeyValue) *resource.Resource {
	opts := []resource.Option{resource.WithAttributes(append(attrs, svc.Attributes()...)...)}
	if res&ResEnv != 0 {
		opts = append(opts, resource.WithFromEnv())
	}
	if res&ResHost != 0 {
		opts = append(opts, resource.WithHost())
	}
	if res&ResTelemetrySDK != 0 {
		opts = append(opts, resource.WithTelemetrySDK())
	}
	if res&ResOS != 0 {
		opts = append(opts, resource.WithOS())
	}
	if res&ResProcess != 0 {
		opts = append(opts, resource.WithProcess())
	}
	if res&ResContainer != 0 {
		opts = append(opts, resource.WithContainer())
	}
	attributes, err := resource.New(ctx, opts...)
	if err != nil {
		return resource.Default()
	}
	return attributes
}
