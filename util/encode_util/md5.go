package encode_util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(src []byte) string {
	m := md5.New()
	m.Write(src)
	b := m.Sum(nil)
	return hex.EncodeToString(b)
}
func Md5EncodeBySalt(src, salt []byte) string {
	m := md5.New()
	m.Write(src)
	m.Write(salt)
	b := m.Sum(nil)
	return hex.EncodeToString(b)
}

func Md5EncodeStrBySalt(src, salt string) string {
	if src == "" {
		return ""
	}
	if salt == "" {
		return Md5Encode([]byte(src))
	}
	return Md5EncodeBySalt([]byte(src), []byte(salt))
}
