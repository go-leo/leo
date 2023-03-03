package stream

// Header is a mapping from keys to values.
type Header map[string][]string

// Len returns the number of items in Header.
func (h Header) Len() int {
	return len(h)
}

// Get obtains the values for a given key.
func (h Header) Get(k string) []string {
	return h[k]
}

// Set sets the value of a given key with a slice of values.
func (h Header) Set(k string, values ...string) Header {
	if len(values) == 0 {
		return h
	}
	h[k] = values
	return h
}

// Delete delete the values for a given key.
func (h Header) Delete(k string) Header {
	delete(h, k)
	return h
}

// Range iterates the header
func (h Header) Range(fn func(key string, values []string)) Header {
	for key, values := range h {
		fn(key, values)
	}
	return h
}

// Append appends the value of a given key with a slice of values.
func (h Header) Append(k string, values ...string) Header {
	if len(values) == 0 {
		return h
	}
	h[k] = append(h[k], values...)
	return h
}
