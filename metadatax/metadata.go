package metadatax

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

type _Metadata map[string][]string

func (md _Metadata) Set(key string, values ...string) {
	if md == nil {
		return
	}
	if len(values) == 0 {
		return
	}
	key = strings.ToLower(key)
	md[key] = values
}

func (md _Metadata) Append(key string, values ...string) {
	if md == nil {
		return
	}
	if len(values) == 0 {
		return
	}
	key = strings.ToLower(key)
	md[key] = append(md[key], values...)
}

func (md _Metadata) Get(key string) string {
	if md == nil {
		return ""
	}
	key = strings.ToLower(key)
	values := md[key]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func (md _Metadata) Values(key string) []string {
	if md == nil {
		return nil
	}
	key = strings.ToLower(key)
	return md[key]
}

func (md _Metadata) Keys() []string {
	if md == nil {
		return nil
	}
	res := make([]string, 0, len(md))
	for key := range md {
		res = append(res, strings.ToLower(key))
	}
	return res
}

func (md _Metadata) Range(f func(key string, values []string) bool) {
	for key, values := range md {
		if !f(strings.ToLower(key), values) {
			break
		}
	}
}

func (md _Metadata) Delete(key string) {
	if md == nil {
		return
	}
	key = strings.ToLower(key)
	delete(md, key)
}

func (md _Metadata) Len() int {
	if md == nil {
		return 0
	}
	return len(md)
}

func (md _Metadata) Clone() Metadata {
	if md == nil {
		return nil
	}
	res := _Metadata{}
	for _, key := range md.Keys() {
		res[key] = slices.Clone(md[key])
	}
	return res
}

func New() Metadata {
	return _Metadata{}
}

func Join(mds ...Metadata) Metadata {
	res := New()
	for _, m := range mds {
		if m == nil || m.Len() <= 0 {
			continue
		}
		m.Range(func(key string, values []string) bool {
			key = strings.ToLower(key)
			res.Append(key, values...)
			return true
		})
	}
	return res
}

func Pairs(kv ...string) Metadata {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("metadatax: Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}
	res := New()
	for i := 0; i < len(kv); i += 2 {
		res.Append(kv[i], kv[i+1])
	}
	return res
}

func FromMap[M ~map[string][]string](m M) Metadata {
	res := New()
	for key, values := range m {
		res.Append(key, values...)
	}
	return res
}
