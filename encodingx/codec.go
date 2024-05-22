package encodingx

import (
	"encoding/json"
	"io"
	"sync"
)

type Encoder interface {
	Encode(val any) error
}

type Decoder interface {
	Decode(obj any) error
}

type JSONEncoder interface {
	Encoder
	SetIndent(prefix, indent string)
	SetEscapeHTML(escapeHTML bool)
}

type EncoderFactory func(w io.Writer) Encoder

type DecoderFactory func(r io.Reader) Decoder

var (
	// defaultEncoderFactories
	defaultEncoderFactories map[string]EncoderFactory

	// defaultDecoderFactories
	defaultDecoderFactories map[string]DecoderFactory

	// encodingMutex
	encodingMutex sync.RWMutex
)

func init() {
	defaultEncoderFactories = make(map[string]EncoderFactory)
	defaultEncoderFactories["json"] = func(w io.Writer) Encoder { return json.NewEncoder(w) }

	defaultDecoderFactories = make(map[string]DecoderFactory)
	defaultDecoderFactories["json"] = func(r io.Reader) Decoder { return json.NewDecoder(r) }
}

func RegisterEncoderFactory(name string, factory EncoderFactory) {
	if factory == nil {
		panic("encodingx: Register EncoderFactory is nil")
	}
	encodingMutex.Lock()
	defaultEncoderFactories[name] = factory
	encodingMutex.Unlock()
}

func RegisterDecoderFactory(name string, factory DecoderFactory) {
	if factory == nil {
		panic("encodingx: Register DecoderFactory is nil")
	}
	encodingMutex.Lock()
	defaultDecoderFactories[name] = factory
	encodingMutex.Unlock()
}

func GetEncoderFactory(name string) EncoderFactory {
	encodingMutex.RUnlock()
	factory := defaultEncoderFactories[name]
	encodingMutex.RUnlock()
	return factory
}

func GetDecoderFactory(name string) DecoderFactory {
	encodingMutex.RUnlock()
	factory := defaultDecoderFactories[name]
	encodingMutex.RUnlock()
	return factory
}
