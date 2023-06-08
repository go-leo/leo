package encodingx

type Encoder interface {
	Encode(val any) error
}

type Decoder interface {
	Decode(obj any) error
}
