package metadatax

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	grpcmetadata "google.golang.org/grpc/metadata"
	"net/http"
)

type metadata map[string][]string

func (m metadata) Set(key string, value ...string) {
	if m == nil {
		return
	}
	m[key] = value
}

func (m metadata) Append(key string, value ...string) {
	if m == nil {
		return
	}
	m[key] = append(m[key], value...)
}

func (m metadata) Get(key string) string {
	if m == nil {
		return ""
	}
	v := m[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func (m metadata) Values(key string) []string {
	if m == nil {
		return nil
	}
	return m[key]
}

func (m metadata) Keys() []string {
	return maps.Keys(m)
}

func (m metadata) Delete(key string) {
	if m == nil {
		return
	}
	delete(m, key)
}

func (m metadata) Len() int {
	return len(m)
}

func (m metadata) Clone() Metadata {
	if m == nil {
		return nil
	}
	clonedMd := metadata{}
	for _, key := range m.Keys() {
		clonedMd[key] = slices.Clone(m[key])
	}
	return clonedMd
}

func New() Metadata {
	return metadata{}
}

func FromHttpHeader(header http.Header) Metadata {
	return metadata(header)
}

func AsHttpHeader(metadata Metadata) http.Header {
	header := http.Header{}
	for _, key := range metadata.Keys() {
		values := metadata.Values(key)
		for _, value := range values {
			header.Add(key, value)
		}
	}
	return header
}

func FromGrpcMetadata(md grpcmetadata.MD) Metadata {
	return metadata(md)
}

func AsGrpcMetadata(metadata Metadata) grpcmetadata.MD {
	md := grpcmetadata.MD{}
	for _, key := range metadata.Keys() {
		md.Set(key, metadata.Values(key)...)
	}
	return md
}
