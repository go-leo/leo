package httpx

import "net/http"

// DisableKeepAlivesClient returns a new http.Client with similar default values to
// http.Client, but with a non-shared Transport, idle connections disabled, and
// keepalives disabled.
func DisableKeepAlivesClient() *http.Client {
	return &http.Client{
		Transport: DisableKeepAlivesTransport(),
	}
}

// PooledClient returns a new http.Client with similar default values to
// http.Client, but with a shared Transport. Do not use this function for
// transient clients as it can leak file descriptors over time. Only use this
// for clients that will be re-used for the same host(s).
func PooledClient() *http.Client {
	return &http.Client{
		Transport: PooledTransport(),
	}
}
