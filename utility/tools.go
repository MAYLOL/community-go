package utility

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256(source string) string {
	h := sha256.New()
	h.Write([]byte(source))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
