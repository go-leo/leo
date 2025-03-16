package consulx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/stainx"
	"google.golang.org/grpc"
	"net"
	"net/url"
	"strconv"
	"strings"
)

var _ sdx.Builder = (*Builder)(nil)

// schemeName for the urls
// All target URLs like 'etcdv3://.../...' will be resolved by this builder
const schemeName = "etcdv3"

type Builder struct {
	// DialOptions is a list of dial options for the gRPC client (e.g., for interceptors).
	// For example, pass grpc.WithBlock() to block until the underlying connection is up.
	// Without this, Dial returns immediately and connecting the server happens in background.
	DialOptions   []grpc.DialOption
	ClientCreator func(ctx context.Context, rawURL *url.URL, color string, dialOptions ...grpc.DialOption) (etcdv3.Client, error)
}

func (Builder) Scheme() string {
	return schemeName
}

func (b Builder) BuildInstancer(ctx context.Context, instance *url.URL, color string, logger kitlog.Logger) (sd.Instancer, error) {
	dsn := strings.Join([]string{schemeName + ":/", instance.Host, instance.Path + "?" + instance.RawQuery}, "/")
	rawURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("malformed url, %v", err)
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", dsn)
	}

	if b.ClientCreator == nil {
		b.ClientCreator = DefaultClientCreator
	}
	client, err := b.ClientCreator(ctx, rawURL, color, b.DialOptions...)
	if err != nil {
		return nil, err
	}

	service := strings.TrimLeft(rawURL.Path, "/")
	color, ok := stainx.ColorExtractor(ctx)
	if !ok {
		prefix := fmt.Sprintf("/leo/services/%s", service)
		return etcdv3.NewInstancer(client, prefix, logger)
	}
	prefix := fmt.Sprintf("/leo/services/%s/%s", service, color)
	return etcdv3.NewInstancer(client, prefix, logger)
}

func (b Builder) BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string, logger kitlog.Logger) (sd.Registrar, error) {
	dsn := strings.Join([]string{schemeName + ":/", instance.Host, instance.Path + "?" + instance.RawQuery}, "/")
	rawURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("malformed url, %v", err)
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", dsn)
	}

	if b.ClientCreator == nil {
		b.ClientCreator = DefaultClientCreator
	}
	client, err := b.ClientCreator(ctx, rawURL, color, b.DialOptions...)
	if err != nil {
		return nil, err
	}

	service := strings.TrimLeft(rawURL.Path, "/")
	prefix := fmt.Sprintf("/leo/services/%s", service)
	if len(color) >= 0 {
		prefix = fmt.Sprintf("/leo/services/%s/%s", prefix, color)
	}
	addr := net.JoinHostPort(ip.String(), strconv.Itoa(port))
	key := fmt.Sprintf("/leo/services/%s/%s", prefix, addr)
	return etcdv3.NewRegistrar(client, etcdv3.Service{Key: key, Value: addr}, logger), nil
}
