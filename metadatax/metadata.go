package metadatax

import (
	"fmt"
	"golang.org/x/exp/slices"
	grpcmetadata "google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

type metadata map[string][]string

func (m metadata) Set(key string, values ...string) {
	if m == nil {
		return
	}
	if len(values) == 0 {
		return
	}
	key = strings.ToLower(key)
	m[key] = values
}

func (m metadata) Append(key string, values ...string) {
	if m == nil {
		return
	}
	if len(values) == 0 {
		return
	}
	key = strings.ToLower(key)
	m[key] = append(m[key], values...)
}

func (m metadata) Get(key string) string {
	if m == nil {
		return ""
	}
	key = strings.ToLower(key)
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
	key = strings.ToLower(key)
	return m[key]
}

func (m metadata) Keys() []string {
	if m == nil {
		return nil
	}
	r := make([]string, 0, len(m))
	for key := range m {
		r = append(r, strings.ToLower(key))
	}
	return r
}

func (m metadata) Range(f func(key string, values []string) bool) {
	for key, values := range m {
		if !f(strings.ToLower(key), values) {
			break
		}
	}
}

func (m metadata) Delete(key string) {
	if m == nil {
		return
	}
	key = strings.ToLower(key)
	delete(m, key)
}

func (m metadata) Len() int {
	if m == nil {
		return 0
	}
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

func FromMap(m map[string][]string) Metadata {
	md := make(metadata, len(m))
	for key, values := range m {
		key = strings.ToLower(key)
		md[key] = values
	}
	return md
}

func Join(mds ...Metadata) Metadata {
	m := metadata{}
	for _, md := range mds {
		if md == nil {
			continue
		}
		md.Range(func(key string, values []string) bool {
			key = strings.ToLower(key)
			m[key] = append(m[key], values...)
			return true
		})
	}
	return m
}

func Pairs(kv ...string) Metadata {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("metadatax: Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}
	md := make(metadata, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		key := strings.ToLower(kv[i])
		value := kv[i+1]
		md.Append(key, value)
	}
	return md
}

// FromGrpcMetadata Convert metadata.MD to Metadata
//
// the key is converted to lowercase.
func FromGrpcMetadata(grpcMD grpcmetadata.MD) Metadata {
	md := New()
	for key, values := range grpcMD {
		md.Set(key, values...)
	}
	return md
}

// FromHttpHeader Convert http.Header to Metadata
//
// The keys should be in canonical form, as returned by http.CanonicalHeaderKey.
func FromHttpHeader(header http.Header) Metadata {
	md := New()
	for key, values := range header {
		md.Set(key, values...)
	}
	return md
}

// AsGrpcMetadata Convert Metadata to metadata.MD
func AsGrpcMetadata(md Metadata) grpcmetadata.MD {
	grpcMD := grpcmetadata.MD{}
	for _, key := range md.Keys() {
		grpcMD.Set(key, md.Values(key)...)
	}
	return grpcMD
}

// AsHttpHeader Convert Metadata to http.Header
func AsHttpHeader(md Metadata) http.Header {
	header := http.Header{}
	for _, key := range md.Keys() {
		values := md.Values(key)
		for _, value := range values {
			header.Add(key, value)
		}
	}
	return header
}
