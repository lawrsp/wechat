package minip

import (
	"crypto/sha1"
)

func CheckDataSignature(sessionKey string, rawData string, signature string) bool {
	h := sha1.New()
	h.Write([]byte(rawData))
	h.Write([]byte(sessionKey))
	check := h.Sum(nil)
	return string(check) == signature
}
