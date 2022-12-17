package registry

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/spf13/cast"
)

// ServiceInfo is used to register a new service.
type ServiceInfo struct {
	// ID is service id.
	ID string
	// Name is service name.
	Name string
	// Transport, "HTTP" or "gRPC".
	Transport string
	// Host is the host that Endpoint used.
	Host string
	// Port is the port that Endpoint used.
	Port int
	// Metadata is other information the service carried.
	Metadata map[string]string
	// Version is version of the service
	Version string
}

func (s *ServiceInfo) String() string {
	if s == nil {
		return ""
	}
	data, _ := json.Marshal(s)
	return string(data)
}

func (s *ServiceInfo) Clone() *ServiceInfo {
	if s == nil {
		return nil
	}
	copy := *s
	copy.Metadata = make(map[string]string, len(s.Metadata))
	for key, val := range s.Metadata {
		copy.Metadata[key] = val
	}
	return &copy
}

func ServiceInfoFromURL(uri url.URL, transport string) *ServiceInfo {
	metaData := uri.Query().Get("metadata")
	md := make(map[string]string)
	_ = json.Unmarshal([]byte(metaData), &md)
	return &ServiceInfo{
		ID:        "",
		Name:      strings.TrimLeft(uri.Path, "/"),
		Transport: transport,
		Host:      uri.Hostname(),
		Port:      cast.ToInt(uri.Port()),
		Metadata:  md,
		Version:   uri.Query().Get("version"),
	}
}
