package stringx

import (
	"strconv"
	"strings"
)

type Builder struct {
	b *strings.Builder
}

func (b *Builder) checkNil() {
	if b.b == nil {
		b.b = &strings.Builder{}
	}
}

// String returns the accumulated string.
func (b *Builder) String() string {
	b.checkNil()
	return b.b.String()
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *Builder) Len() int {
	b.checkNil()
	return b.b.Len()
}

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder) Cap() int {
	b.checkNil()
	return b.b.Cap()
}

// Reset resets the Builder to be empty.
func (b *Builder) Reset() {
	b.checkNil()
	b.b.Reset()
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (b *Builder) Grow(n int) {
	b.checkNil()
	b.b.Grow(n)
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error) {
	b.checkNil()
	return b.b.Write(p)
}

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (b *Builder) WriteByte(c byte) error {
	b.checkNil()
	return b.b.WriteByte(c)
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
func (b *Builder) WriteRune(r rune) (int, error) {
	b.checkNil()
	return b.b.WriteRune(r)
}

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *Builder) WriteString(s string) (int, error) {
	b.checkNil()
	return b.b.WriteString(s)
}

// WriteInt appends the string form of the integer i.
func (b *Builder) WriteInt(i int64, base int) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendInt(nil, i, base))
	return err
}

// WriteUint appends the string form of the unsigned integer i.
func (b *Builder) WriteUint(i uint64, base int) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendUint(nil, i, base))
	return err
}

// WriteBool appends "true" or "false".
func (b *Builder) WriteBool(bl bool) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendBool(nil, bl))
	return err
}

// WriteFloat appends the string form of the floating-point number f.
func (b *Builder) WriteFloat(f float64, fmt byte, prec, bitSize int) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendFloat(nil, f, fmt, prec, bitSize))
	return err
}

// WriteQuote appends a double-quoted Go string literal representing s.
func (b *Builder) WriteQuote(s string) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuote(nil, s))
	return err
}

// WriteQuoteRune appends a single-quoted Go character literal representing the rune.
func (b *Builder) WriteQuoteRune(r rune) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteRune(nil, r))
	return err
}

// WriteQuoteRuneToASCII appends a single-quoted Go character literal representing the rune.
func (b *Builder) WriteQuoteRuneToASCII(r rune) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteRuneToASCII(nil, r))
	return err
}

// WriteQuoteRuneToGraphic appends a single-quoted Go character literal representing the rune.
func (b *Builder) WriteQuoteRuneToGraphic(r rune) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteRuneToGraphic(nil, r))
	return err
}

// WriteQuoteToASCII appends a double-quoted Go string literal representing s.
func (b *Builder) WriteQuoteToASCII(s string) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteToASCII(nil, s))
	return err
}

// WriteQuoteToGraphic appends a double-quoted Go string literal representing s.
func (b *Builder) WriteQuoteToGraphic(s string) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteToGraphic(nil, s))
	return err
}

func NewBuilder() *Builder {
	return &Builder{b: &strings.Builder{}}
}

func NewBuilderBuilder(b *strings.Builder) *Builder {
	return &Builder{b: b}
}
