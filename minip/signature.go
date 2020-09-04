package minip

import (
	"bytes"
	"crypto/sha1"
)

func CheckDataSignature(sessionKey string, rawData []byte, signature []byte) bool {
	h := sha1.New()
	h.Write([]byte(rawData))
	h.Write([]byte(sessionKey))
	check := h.Sum(nil)

	return bytes.Equal(check, signature)
}
