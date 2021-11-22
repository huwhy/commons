package encode_util

import "testing"

func TestMd5Encode(t *testing.T) {
	s := Md5EncodeStrBySalt("abc123", "123")
	t.Log(s)
	s = Md5EncodeStrBySalt("abc123", "123")
	t.Log(s)
}
