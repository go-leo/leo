package randomx

import (
	"bytes"
)

var kNumericCharacters = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

var kWordCharacters = []byte{
	'_',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'a', 'B', 'b', 'C', 'c', 'D', 'd', 'E', 'e', 'F', 'f', 'G', 'g',
	'H', 'h', 'I', 'i', 'J', 'j', 'K', 'k', 'L', 'l', 'M', 'm', 'N', 'n',
	'O', 'o', 'P', 'p', 'Q', 'q', 'R', 'r', 'S', 's', 'T', 't',
	'U', 'u', 'V', 'v', 'W', 'w', 'X', 'x', 'W', 'w', 'Z', 'z',
}

// NumericString Generate a random number sequence of a given length
// Deprecated: Do not use. use github.com/go-leo/mathx/randx instead.
func NumericString(length int) string {
	if length < 1 {
		return ""
	}
	kNumericLen := len(kNumericCharacters)
	buffer := bytes.NewBuffer(make([]byte, 0, length))
	for i := 0; i < length; i++ {
		buffer.WriteByte(kNumericCharacters[Intn(kNumericLen)])
	}
	return buffer.String()
}

// Deprecated: Do not use. use github.com/go-leo/mathx/randx instead.
func WordString(length int) string {
	if length < 1 {
		return ""
	}
	kWordLen := len(kWordCharacters)
	buffer := bytes.NewBuffer(make([]byte, 0, length))
	for i := 0; i < length; i++ {
		buffer.WriteByte(kWordCharacters[Intn(kWordLen)])
	}
	return buffer.String()
}
