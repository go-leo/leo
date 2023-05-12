package message

// A Header represents the key-value pairs in a header.
type Header map[string][]string

// Add adds the key, value pair to the header.
func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

// Set sets the header entries associated with key to the
// single element value. It replaces any existing values
// associated with key.
func (h Header) Set(key, value string) {
	h[key] = []string{value}
}

// Get gets the first value associated with the given key. If
// there are no values associated with the key, Get returns "".
func (h Header) Get(key string) string {
	if h == nil {
		return ""
	}
	v := h[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

// Values returns all values associated with the given key.
func (h Header) Values(key string) []string {
	if h == nil {
		return nil
	}
	return h[key]
}

// Del deletes the values associated with key.
func (h Header) Del(key string) {
	delete(h, key)
}

// Clone returns a copy of h or nil if h is nil.
func (h Header) Clone() Header {
	if h == nil {
		return nil
	}

	// Find total number of values.
	nv := 0
	for _, vv := range h {
		nv += len(vv)
	}
	sv := make([]string, nv) // shared backing array for headers' values
	h2 := make(Header, len(h))
	for k, vv := range h {
		if vv == nil {
			// Preserve nil values. ReverseProxy distinguishes
			// between nil and zero-length header values.
			h2[k] = nil
			continue
		}
		n := copy(sv, vv)
		h2[k] = sv[:n:n]
		sv = sv[n:]
	}
	return h2
}
