package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func md5String(src string) string {
	byteString := md5.Sum([]byte(src))
	return hex.EncodeToString(byteString[:])
}
