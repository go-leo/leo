package internal

import (
	"net/url"
	"strings"
)

func ExtractEndpoint(instance *url.URL) string {
	endpoint := instance.Path
	if endpoint == "" {
		endpoint = instance.Opaque
	}
	// For targets of the form "[scheme]://[authority]/endpoint, the endpoint
	// value returned from url.Parse() contains a leading "/". Although this is
	// in accordance with RFC 3986, we do not want to break existing resolver
	// implementations which expect the endpoint without the leading "/". So, we
	// end up stripping the leading "/" here. But this will result in an
	// incorrect parsing for something like "unix:///path/to/socket". Since we
	// own the "unix" resolver, we can workaround in the unix resolver by using
	// the `URL` field.
	return strings.TrimPrefix(endpoint, "/")
}
