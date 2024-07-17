package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5SaltHash(text, salt string) string {
	return MD5Hash(text + ":" + salt)
}

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
