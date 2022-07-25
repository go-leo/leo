package bufx

import (
	"io"
	"strconv"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() any {
		return &ByteBuffer{B: make([]byte, 0)}
	},
}

func Get() *ByteBuffer {
	v := pool.Get()
	if v != nil {
		return v.(*ByteBuffer)
	}
	return &ByteBuffer{B: make([]byte, 0)}
}

type ByteBuffer struct {
	B []byte
}

func (b *ByteBuffer) Len() int {
	return len(b.B)
}

func (b *ByteBuffer) Cap() int {
	return cap(b.B)
}

func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
	p := b.B
	nStart := int64(len(p))
	nMax := int64(cap(p))
	n := nStart
	if nMax == 0 {
		nMax = 64
		p = make([]byte, nMax)
	} else {
		p = p[:nMax]
	}
	for {
		if n == nMax {
			nMax *= 2
			bNew := make([]byte, nMax)
			copy(bNew, p)
			p = bNew
		}
		nn, err := r.Read(p[n:])
		n += int64(nn)
		if err != nil {
			b.B = p[:n]
			n -= nStart
			if err == io.EOF {
				return n, nil
			}
			return n, err
		}
	}
}

func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.B)
	return int64(n), err
}

func (b *ByteBuffer) Write(p []byte) (int, error) {
	b.B = append(b.B, p...)
	return len(p), nil
}

func (b *ByteBuffer) WriteByte(c byte) error {
	b.B = append(b.B, c)
	return nil
}

func (b *ByteBuffer) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

func (b *ByteBuffer) Set(p []byte) {
	b.B = append(b.B[:0], p...)
}

func (b *ByteBuffer) SetString(s string) {
	b.B = append(b.B[:0], s...)
}

func (b *ByteBuffer) AppendInt(i int64) {
	b.B = strconv.AppendInt(b.B, i, 10)
}

func (b *ByteBuffer) AppendTime(t time.Time, layout string) {
	b.B = t.AppendFormat(b.B, layout)
}

func (b *ByteBuffer) AppendUint(i uint64) {
	b.B = strconv.AppendUint(b.B, i, 10)
}

func (b *ByteBuffer) AppendBool(v bool) {
	b.B = strconv.AppendBool(b.B, v)
}

func (b *ByteBuffer) AppendFloat(f float64, bitSize int) {
	b.B = strconv.AppendFloat(b.B, f, 'f', -1, bitSize)
}

func (b *ByteBuffer) String() string {
	return string(b.B)
}

func (b *ByteBuffer) Bytes() []byte {
	return b.B
}

func (b *ByteBuffer) Reset() {
	b.B = b.B[:0]
}

func (b *ByteBuffer) Free() {
	b.Reset()
	pool.Put(b)
}
