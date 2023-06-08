package formx

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
	"errors"
	"net/url"
	"sync"
)

var defaultEncoder = form.NewEncoder()
var encoders = map[string]*form.Encoder{}
var encoderMu sync.RWMutex

func init() {
	//defaultEncoder.SetNamespacePrefix()
}

func RegisterEncoder(tag string, encoderFunc func() *form.Encoder) {
	encoderMu.Lock()
	defer encoderMu.Unlock()
	encoders[tag] = encoderFunc()
}

var defaultDecoder = form.NewDecoder()
var decoders = map[string]*form.Decoder{}
var decoderMu sync.RWMutex

func RegisterDecoder(tag string, decoderFunc func() *form.Decoder) {
	decoderMu.Lock()
	defer decoderMu.Unlock()
	decoders[tag] = decoderFunc()
}

func Marshal(v any, tag ...string) (url.Values, error) {
	encoder, err := getEncoder(tag)
	if err != nil {
		return nil, err
	}
	return encoder.Encode(v)
}

func Unmarshal(form url.Values, v any, tag ...string) error {
	decoder, err := getDecoder(tag)
	if err != nil {
		return err
	}

	return decoder.Decode(v, form)
}

func NewEncoder(form url.Values, tag ...string) encodingx.Encoder {
	return &encoder{form: form, tag: tag}
}

func NewDecoder(form url.Values, tag ...string) encodingx.Decoder {
	return &decoder{form: form, tag: tag}
}

type encoder struct {
	tag  []string
	form url.Values
}

func (e *encoder) Encode(val any) error {
	data, err := Marshal(val, e.tag...)
	if err != nil {
		return err
	}
	for key, vals := range data {
		for _, val := range vals {
			e.form.Add(key, val)
		}
	}
	return nil
}

type decoder struct {
	tag  []string
	form url.Values
}

func (d *decoder) Decode(obj any) error {
	return Unmarshal(d.form, obj, d.tag...)
}

func getEncoder(tag []string) (*form.Encoder, error) {
	var encoder *form.Encoder
	if slicex.IsEmpty(tag) {
		return defaultEncoder, nil
	}
	encoderMu.RLock()
	encoder, ok := encoders[tag[0]]
	encoderMu.RUnlock()
	if !ok {
		return nil, errors.New("not found encoder for tag: " + tag[0])
	}
	return encoder, nil
}

func getDecoder(tag []string) (*form.Decoder, error) {
	if slicex.IsEmpty(tag) {
		return defaultDecoder, nil
	}
	decoderMu.RLock()
	decoder, ok := decoders[tag[0]]
	decoderMu.RUnlock()
	if !ok {
		return nil, errors.New("not found decoder for tag: " + tag[0])
	}
	return decoder, nil
}
