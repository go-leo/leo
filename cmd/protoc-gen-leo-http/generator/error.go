package generator

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"strings"
)

func errInvalidType(endpoint *internal.Endpoint, names []string) error {
	return fmt.Errorf("%s, %s field type invalid", endpoint.FullName(), strings.Join(names, "."))
}

func errNotFoundField(endpoint *internal.Endpoint, names []string) error {
	return fmt.Errorf("%s, failed to find field %s", endpoint.FullName(), strings.Join(names, "."))
}
