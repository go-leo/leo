package bytex

import (
	"bytes"
	"io"
	"strconv"
)

type Buffer struct {
	b *bytes.Buffer
}

func (b *Buffer) checkNil() {
	if b.b == nil {
		b.b = &bytes.Buffer{}
	}
}

// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
// The slice is valid for use only until the next buffer modification (that is,
// only until the next call to a method like Read, Write, Reset, or Truncate).
// The slice aliases the buffer content at least until the next buffer modification,
// so immediate changes to the slice will affect the result of future reads.
func (b *Buffer) Bytes() []byte {
	b.checkNil()
	return b.b.Bytes()
}

// String returns the contents of the unread portion of the buffer
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//
// To build strings more efficiently, see the strings.Builder type.
func (b *Buffer) String() string {
	b.checkNil()
	return b.b.String()
}

// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Bytes()).
func (b *Buffer) Len() int {
	b.checkNil()
	return b.b.Len()
}

// Cap returns the capacity of the buffer's underlying byte slice, that is, the
// total space allocated for the buffer's data.
func (b *Buffer) Cap() int {
	b.checkNil()
	return b.b.Cap()
}

// Truncate discards all but the first n unread bytes from the buffer
// but continues to use the same allocated storage.
// It panics if n is negative or greater than the length of the buffer.
func (b *Buffer) Truncate(n int) {
	b.checkNil()
	b.b.Truncate(n)
}

// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as Truncate(0).
func (b *Buffer) Reset() {
	b.checkNil()
	b.b.Reset()
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to the
// buffer without another allocation.
// If n is negative, Grow will panic.
// If the buffer can't grow it will panic with ErrTooLarge.
func (b *Buffer) Grow(n int) {
	b.checkNil()
	b.b.Grow(n)
}

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.checkNil()
	return b.b.Write(p)
}

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with ErrTooLarge.
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.checkNil()
	return b.b.WriteString(s)
}

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	b.checkNil()
	return b.b.ReadFrom(r)
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of bytes written; it always fits into an
// int, but it is int64 to match the io.WriterTo interface. Any error
// encountered during the write is also returned.
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	b.checkNil()
	return b.b.WriteTo(b)
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// The returned error is always nil, but is included to match bufio.Writer's
// WriteByte. If the buffer becomes too large, WriteByte will panic with
// ErrTooLarge.
func (b *Buffer) WriteByte(c byte) error {
	b.checkNil()
	return b.b.WriteByte(c)
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to the
// buffer, returning its length and an error, which is always nil but is
// included to match bufio.Writer's WriteRune. The buffer is grown as needed;
// if it becomes too large, WriteRune will panic with ErrTooLarge.
func (b *Buffer) WriteRune(r rune) (n int, err error) {
	b.checkNil()
	return b.b.WriteRune(r)
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
func (b *Buffer) Read(p []byte) (n int, err error) {
	b.checkNil()
	return b.b.Read(p)
}

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The slice is only valid until the next call to a read or write method.
func (b *Buffer) Next(n int) []byte {
	b.checkNil()
	return b.b.Next(n)
}

// ReadByte reads and returns the next byte from the buffer.
// If no byte is available, it returns error io.EOF.
func (b *Buffer) ReadByte() (byte, error) {
	b.checkNil()
	return b.b.ReadByte()
}

// ReadRune reads and returns the next UTF-8-encoded
// Unicode code point from the buffer.
// If no bytes are available, the error returned is io.EOF.
// If the bytes are an erroneous UTF-8 encoding, it
// consumes one byte and returns U+FFFD, 1.
func (b *Buffer) ReadRune() (r rune, size int, err error) {
	b.checkNil()
	return b.b.ReadRune()
}

// UnreadRune unreads the last rune returned by ReadRune.
// If the most recent read or write operation on the buffer was
// not a successful ReadRune, UnreadRune returns an error.  (In this regard
// it is stricter than UnreadByte, which will unread the last byte
// from any read operation.)
func (b *Buffer) UnreadRune() error {
	b.checkNil()
	return b.b.UnreadRune()
}

// UnreadByte unreads the last byte returned by the most recent successful
// read operation that read at least one byte. If a write has happened since
// the last read, if the last read returned an error, or if the read read zero
// bytes, UnreadByte returns an error.
func (b *Buffer) UnreadByte() error {
	b.checkNil()
	return b.b.UnreadByte()
}

// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadBytes returns err != nil if and only if the returned data does not end in
// delim.
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	b.checkNil()
	return b.b.ReadBytes(delim)
}

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end
// in delim.
func (b *Buffer) ReadString(delim byte) (line string, err error) {
	b.checkNil()
	return b.b.ReadString(delim)
}

// NewBuffer creates and initializes a new Buffer using buf as its
// initial contents. The new Buffer takes ownership of buf, and the
// caller should not use buf after this call. NewBuffer is intended to
// prepare a Buffer to read existing data. It can also be used to set
// the initial size of the internal buffer for writing. To do that,
// buf should have the desired capacity but a length of zero.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.
func NewBuffer(buf []byte) *Buffer { return &Buffer{b: bytes.NewBuffer(buf)} }

// NewBufferString creates and initializes a new Buffer using string s as its
// initial contents. It is intended to prepare a buffer to read an existing
// string.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.
func NewBufferString(s string) *Buffer {
	return &Buffer{b: bytes.NewBufferString(s)}
}

// WriteInt appends the string form of the integer i.
func (b *Buffer) WriteInt(i int64, base int) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendInt(nil, i, base))
	return err
}

// WriteUint appends the string form of the unsigned integer i.
func (b *Buffer) WriteUint(i uint64, base int) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendUint(nil, i, base))
	return err
}

// WriteBool appends "true" or "false".
func (b *Buffer) WriteBool(bl bool) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendBool(nil, bl))
	return err
}

// WriteFloat appends the string form of the floating-point number f.
func (b *Buffer) WriteFloat(f float64, fmt byte, prec, bitSize int) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendFloat(nil, f, fmt, prec, bitSize))
	return err
}

// WriteQuote appends a double-quoted Go string literal representing s.
func (b *Buffer) WriteQuote(s string) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuote(nil, s))
	return err
}

// WriteQuoteRune appends a single-quoted Go character literal representing the rune.
func (b *Buffer) WriteQuoteRune(r rune) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteRune(nil, r))
	return err
}

// WriteQuoteRuneToASCII appends a single-quoted Go character literal representing the rune.
func (b *Buffer) WriteQuoteRuneToASCII(r rune) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteRuneToASCII(nil, r))
	return err
}

// WriteQuoteRuneToGraphic appends a single-quoted Go character literal representing the rune.
func (b *Buffer) WriteQuoteRuneToGraphic(r rune) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteRuneToGraphic(nil, r))
	return err
}

// WriteQuoteToASCII appends a double-quoted Go string literal representing s.
func (b *Buffer) WriteQuoteToASCII(s string) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteToASCII(nil, s))
	return err
}

// WriteQuoteToGraphic appends a double-quoted Go string literal representing s.
func (b *Buffer) WriteQuoteToGraphic(s string) error {
	b.checkNil()
	_, err := b.b.Write(strconv.AppendQuoteToGraphic(nil, s))
	return err
}

func NewBufferBuffer(b *bytes.Buffer) *Buffer {
	return &Buffer{b: b}
}
