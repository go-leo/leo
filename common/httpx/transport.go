package httpx

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

type TransportBuilder struct {
	proxy                  func(*http.Request) (*url.URL, error)
	dial                   func(ctx context.Context, network string, addr string) (net.Conn, error)
	dialTLS                func(ctx context.Context, network string, addr string) (net.Conn, error)
	tlsConfig              *tls.Config
	tlsHandshakeTimeout    time.Duration
	disableKeepAlives      bool
	disableCompression     bool
	maxIdleConns           int
	maxIdleConnsPerHost    int
	maxConnsPerHost        int
	idleConnTimeout        time.Duration
	responseHeaderTimeout  time.Duration
	expectContinueTimeout  time.Duration
	tlsNextProto           map[string]func(authority string, c *tls.Conn) http.RoundTripper
	proxyConnectHeader     http.Header
	getProxyConnectHeader  func(ctx context.Context, proxyURL *url.URL, target string) (http.Header, error)
	maxResponseHeaderBytes int64
	writeBufferSize        int
	readBufferSize         int
	forceAttemptHTTP2      bool
}

func (builder *TransportBuilder) Proxy(proxy func(*http.Request) (*url.URL, error)) *TransportBuilder {
	builder.proxy = proxy
	return builder
}

func (builder *TransportBuilder) Dial(dial func(ctx context.Context, network string, addr string) (net.Conn, error)) *TransportBuilder {
	builder.dial = dial
	return builder
}

func (builder *TransportBuilder) DialTLS(dialTLS func(ctx context.Context, network string, addr string) (net.Conn, error)) *TransportBuilder {
	builder.dialTLS = dialTLS
	return builder
}

func (builder *TransportBuilder) TLSConfig(tlsConfig *tls.Config) *TransportBuilder {
	builder.tlsConfig = tlsConfig
	return builder
}

func (builder *TransportBuilder) TLSHandshakeTimeout(timeout time.Duration) *TransportBuilder {
	builder.tlsHandshakeTimeout = timeout
	return builder
}

func (builder *TransportBuilder) DisableKeepAlives(disable bool) *TransportBuilder {
	builder.disableKeepAlives = disable
	return builder
}

func (builder *TransportBuilder) DisableCompression(disable bool) *TransportBuilder {
	builder.disableCompression = disable
	return builder
}

func (builder *TransportBuilder) MaxIdleConns(n int) *TransportBuilder {
	builder.maxIdleConns = n
	return builder
}

func (builder *TransportBuilder) MaxIdleConnsPerHost(n int) *TransportBuilder {
	builder.maxIdleConnsPerHost = n
	return builder
}

func (builder *TransportBuilder) MaxConnsPerHost(n int) *TransportBuilder {
	builder.maxConnsPerHost = n
	return builder
}

func (builder *TransportBuilder) IdleConnTimeout(timeout time.Duration) *TransportBuilder {
	builder.idleConnTimeout = timeout
	return builder
}

func (builder *TransportBuilder) ResponseHeaderTimeout(timeout time.Duration) *TransportBuilder {
	builder.responseHeaderTimeout = timeout
	return builder
}

func (builder *TransportBuilder) ExpectContinueTimeout(timeout time.Duration) *TransportBuilder {
	builder.expectContinueTimeout = timeout
	return builder
}

func (builder *TransportBuilder) TLSNextProto(f map[string]func(authority string, c *tls.Conn) http.RoundTripper) *TransportBuilder {
	builder.tlsNextProto = f
	return builder
}

func (builder *TransportBuilder) ProxyConnectHeader(h http.Header) *TransportBuilder {
	builder.proxyConnectHeader = h
	return builder
}

func (builder *TransportBuilder) GetProxyConnectHeader(f func(ctx context.Context, proxyURL *url.URL, target string) (http.Header, error)) *TransportBuilder {
	builder.getProxyConnectHeader = f
	return builder
}

func (builder *TransportBuilder) MaxResponseHeaderBytes(n int64) *TransportBuilder {
	builder.maxResponseHeaderBytes = n
	return builder
}

func (builder *TransportBuilder) WriteBufferSize(n int) *TransportBuilder {
	builder.writeBufferSize = n
	return builder
}

func (builder *TransportBuilder) ReadBufferSize(n int) *TransportBuilder {
	builder.readBufferSize = n
	return builder
}

func (builder *TransportBuilder) ForceAttemptHTTP2(enable bool) *TransportBuilder {
	builder.forceAttemptHTTP2 = enable
	return builder
}

func (builder *TransportBuilder) Build() *http.Transport {
	return &http.Transport{
		Proxy:                  builder.proxy,
		DialContext:            builder.dial,
		DialTLSContext:         builder.dialTLS,
		TLSClientConfig:        builder.tlsConfig,
		TLSHandshakeTimeout:    builder.tlsHandshakeTimeout,
		DisableKeepAlives:      builder.disableKeepAlives,
		DisableCompression:     builder.disableCompression,
		MaxIdleConns:           builder.maxIdleConns,
		MaxIdleConnsPerHost:    builder.maxIdleConnsPerHost,
		MaxConnsPerHost:        builder.maxConnsPerHost,
		IdleConnTimeout:        builder.idleConnTimeout,
		ResponseHeaderTimeout:  builder.responseHeaderTimeout,
		ExpectContinueTimeout:  builder.expectContinueTimeout,
		TLSNextProto:           builder.tlsNextProto,
		ProxyConnectHeader:     builder.proxyConnectHeader,
		GetProxyConnectHeader:  builder.getProxyConnectHeader,
		MaxResponseHeaderBytes: builder.maxResponseHeaderBytes,
		WriteBufferSize:        builder.writeBufferSize,
		ReadBufferSize:         builder.readBufferSize,
		ForceAttemptHTTP2:      builder.forceAttemptHTTP2,
	}
}

// DisableKeepAlivesTransport returns a new http.Transport with similar default values to
// http.DefaultTransport, but with idle connections and keepalives disabled.
func DisableKeepAlivesTransport() *http.Transport {
	return new(TransportBuilder).
		Proxy(http.ProxyFromEnvironment).
		Dial((&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext).
		ForceAttemptHTTP2(true).
		MaxIdleConns(100).
		IdleConnTimeout(90 * time.Second).
		TLSHandshakeTimeout(10 * time.Second).
		ExpectContinueTimeout(time.Second).
		DisableKeepAlives(true).
		MaxIdleConns(-1).Build()
}

// PooledTransport returns a new http.Transport with similar default
// values to http.DefaultTransport. Do not use this for transient transports as
// it can leak file descriptors over time. Only use this for transports that
// will be re-used for the same host(s).
func PooledTransport() *http.Transport {
	return new(TransportBuilder).
		Proxy(http.ProxyFromEnvironment).
		Dial((&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext).
		ForceAttemptHTTP2(true).
		MaxIdleConns(100).
		IdleConnTimeout(90 * time.Second).
		TLSHandshakeTimeout(10 * time.Second).
		ExpectContinueTimeout(time.Second).
		MaxIdleConnsPerHost(runtime.GOMAXPROCS(0) + 1).Build()
}
