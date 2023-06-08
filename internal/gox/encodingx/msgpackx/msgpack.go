package msgpackx

import (
	"io"

	"github.com/ugorji/go/codec"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	var out []byte
	err := codec.NewEncoderBytes(&out, &codec.MsgpackHandle{}).Encode(v)
	return out, err
}

func Unmarshal(data []byte, v any) error {
	return codec.NewDecoderBytes(data, &codec.MsgpackHandle{}).Decode(v)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return codec.NewEncoder(w, &codec.MsgpackHandle{})
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return codec.NewDecoder(r, &codec.MsgpackHandle{})
}
