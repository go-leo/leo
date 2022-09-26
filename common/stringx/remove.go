package stringx

import "strings"

// Remove takes a string candidate and a string of chars to remove from the candidate.
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func Remove(s string, chars string) string {
	buf := &strings.Builder{}
	for len(s) > 0 {
		idx := strings.IndexAny(s, chars)
		if idx < 0 {
			if buf.Len() > 0 {
				buf.WriteString(s)
			}
			break
		}

		buf.WriteString(s[:idx])
		s = s[idx+1:]
	}

	if buf.Len() == 0 {
		return s
	}
	return buf.String()
}
