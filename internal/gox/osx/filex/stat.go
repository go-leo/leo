package filex

import (
	"os"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
)

// IsDir reports whether the named file is a directory.
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// IsDirectory reports whether the named file is a directory.
var IsDirectory = IsDir

// IsExist returns a boolean indicating whether a file or directory exist.
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// IsNotExist returns a boolean indicating whether a file or directory not exist.
func IsNotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return false
	}
	return os.IsNotExist(err)
}

const (
	Byte         int64 = 1
	Kilobyte           = 1024 * Byte
	Megabyte           = 1024 * Kilobyte
	Gigabyte           = 1024 * Megabyte
	Trillionbyte       = 1024 * Gigabyte
	Petabyte           = 1024 * Trillionbyte
	Exabyte            = 1024 * Petabyte
	// Zettabyte          = 1024 * Exabyte
	// Yottabyte          = 1024 * Zettabyte
	// Brontobyte         = 1024 * Yottabyte
)

func HumanReadableSize(size int64) string {
	s := size
	if s < 0 {
		s = -s
	}
	builder := stringx.Builder{}
	if s >= Exabyte {
		eb := s / Exabyte
		s = s % Exabyte
		_ = builder.WriteInt(eb, 10)
		_, _ = builder.WriteString("EB")
	}

	if s >= Petabyte {
		pb := s / Petabyte
		s = s % Petabyte
		_ = builder.WriteInt(pb, 10)
		_, _ = builder.WriteString("PB")
	}

	if s >= Trillionbyte {
		tb := s / Trillionbyte
		s = s % Trillionbyte
		_ = builder.WriteInt(tb, 10)
		_, _ = builder.WriteString("TB")
	}

	if s >= Gigabyte {
		gb := s / Gigabyte
		s = s % Gigabyte
		_ = builder.WriteInt(gb, 10)
		_, _ = builder.WriteString("GB")
	}

	if s >= Megabyte {
		mb := s / Megabyte
		s = s % Megabyte
		_ = builder.WriteInt(mb, 10)
		_, _ = builder.WriteString("MB")
	}

	if s >= Kilobyte {
		kb := s / Kilobyte
		s = s % Kilobyte
		_ = builder.WriteInt(kb, 10)
		_, _ = builder.WriteString("KB")
	}

	if s >= Byte {
		b := s / Byte
		// s = s % Byte
		_ = builder.WriteInt(b, 10)
		_, _ = builder.WriteString("B")
	}
	return builder.String()
}
