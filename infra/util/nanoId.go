package util

import "crypto/rand"

var (
	defaultAlphabet = []byte("0123456789abcdef")
)

func HexId(size int) string {
	s := make([]byte, size)
	rand.Read(s)
	for i := range s {
		s[i] = defaultAlphabet[s[i]&15]
	}

	return string(s)
}
