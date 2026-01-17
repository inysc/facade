package util

import (
	"bytes"
	"net/http"
)

func FmtHeader(h http.Header) string {
	buf := &bytes.Buffer{}

	buf.WriteByte('{')
	for k, vs := range h {
		if buf.Len() > 1 {
			buf.WriteByte(',')
			buf.WriteByte(' ')
		}
		for i, v := range vs {
			if i > 0 {
				buf.WriteByte(',')
				buf.WriteByte(' ')
			}
			buf.WriteString(k)
			buf.WriteByte(':')
			buf.WriteByte(' ')
			buf.WriteString(v)
		}
	}
	buf.WriteByte('}')

	return buf.String()
}
