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

func Md5EncodeS(src string) string {
	if src == "" {
		return ""
	}
	return Md5Encode([]byte(src))
}
