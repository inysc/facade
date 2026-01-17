package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

func HmacSha512(key, p []byte) string {
	h := hmac.New(sha512.New, key)
	h.Write(p)

	return hex.EncodeToString(h.Sum(nil))
}

func Md5(p string) string {
	h := md5.New()
	h.Write([]byte(p))
	return hex.EncodeToString(h.Sum(nil))
}
