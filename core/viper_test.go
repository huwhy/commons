package core

import (
	"testing"
)

func TestLoadConf(t *testing.T) {
	v, c := LoadConfig()
	t.Log(v, c)
}
