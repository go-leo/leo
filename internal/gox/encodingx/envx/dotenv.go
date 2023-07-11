package envx

import (
	"errors"
	"io"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/convx"

	"github.com/joho/godotenv"
	"golang.org/x/exp/maps"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx"
)

func Marshal(val any) ([]byte, error) {
	envMap, ok := val.(map[string]string)
	if !ok {
		return nil, errors.New("any not convert to map[string]string")
	}
	data, err := godotenv.Marshal(envMap)
	if err != nil {
		return nil, err
	}
	return convx.StringToBytes(data), nil
}

func Unmarshal(data []byte, val any) error {
	envMap, ok := val.(map[string]string)
	if !ok {
		return errors.New("any not convert to map[string]string")
	}
	m, err := godotenv.UnmarshalBytes(data)
	if err != nil {
		return err
	}
	maps.Copy(envMap, m)
	return nil
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return &encoder{w: w}
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return &decoder{r: r}
}

type encoder struct {
	w io.Writer
}

func (e *encoder) Encode(val any) error {
	envMap, ok := val.(map[string]string)
	if !ok {
		return errors.New("any not convert to map[string]string")
	}
	data, err := godotenv.Marshal(envMap)
	if err != nil {
		return err
	}
	_, err = e.w.Write(convx.StringToBytes(data))
	return err
}

type decoder struct {
	r io.Reader
}

func (d *decoder) Decode(obj any) error {
	m, ok := obj.(map[string]string)
	if !ok {
		return errors.New("any not convert to map[string]string")
	}
	envMap, err := godotenv.Parse(d.r)
	if err != nil {
		return err
	}
	maps.Copy(m, envMap)
	return nil
}
