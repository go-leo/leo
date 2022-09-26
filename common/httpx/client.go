package httpx

import (
	"net/http"
	"time"
)

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
type ClientBuilder struct {
	transport     http.RoundTripper
	checkRedirect func(req *http.Request, via []*http.Request) error
	jar           http.CookieJar
	timeout       time.Duration
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func (builder *ClientBuilder) Transport(transport http.RoundTripper) *ClientBuilder {
	builder.transport = transport
	return builder
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func (builder *ClientBuilder) CheckRedirect(f func(req *http.Request, via []*http.Request) error) *ClientBuilder {
	builder.checkRedirect = f
	return builder
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func (builder *ClientBuilder) Jar(jar http.CookieJar) *ClientBuilder {
	builder.jar = jar
	return builder
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func (builder *ClientBuilder) Timeout(timeout time.Duration) *ClientBuilder {
	builder.timeout = timeout
	return builder
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func (builder *ClientBuilder) Build() *http.Client {
	return &http.Client{
		Transport:     builder.transport,
		CheckRedirect: builder.checkRedirect,
		Jar:           builder.jar,
		Timeout:       builder.timeout,
	}
}

// DisableKeepAlivesClient returns a new http.Client with similar default values to
// http.Client, but with a non-shared Transport, idle connections disabled, and
// keepalives disabled.
// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func DisableKeepAlivesClient() *http.Client {
	return new(ClientBuilder).Transport(DisableKeepAlivesTransport()).Build()
}

// PooledClient returns a new http.Client with similar default values to
// http.Client, but with a shared Transport. Do not use this function for
// transient clients as it can leak file descriptors over time. Only use this
// for clients that will be re-used for the same host(s).
// Deprecated: Do not use. use github.com/go-leo/netx/httpx/ClientBuilder instead.
func PooledClient() *http.Client {
	return new(ClientBuilder).Transport(PooledTransport()).Build()
}
