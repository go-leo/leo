package health

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-leo/netx/httpx"
)

type HTTPProber struct {
	timeout         time.Duration
	TLSClientConfig *tls.Config
	httpClient      *http.Client
}

func NewHTTPProber(timeout time.Duration, tlsConfig *tls.Config) *HTTPProber {
	trans := httpx.DisableKeepAlivesTransport()
	trans.TLSClientConfig = tlsConfig
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: trans,
	}
	return &HTTPProber{
		timeout:    timeout,
		httpClient: httpClient,
	}
}

func (probe *HTTPProber) Check(target string) error {
	method := http.MethodGet
	ctx, cancel := context.WithTimeout(context.Background(), probe.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, target, bytes.NewReader(nil))
	if err != nil {
		return err
	}
	resp, err := probe.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	output := new(bytes.Buffer)
	if _, err := io.Copy(output, resp.Body); err != nil {
		return err
	}
	return fmt.Errorf("HTTP %s %s, StatusCode: %d, Status: %s, Output: %s", method, target, resp.StatusCode, resp.Status, output.String())
}
