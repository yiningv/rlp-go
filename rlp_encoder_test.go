package rlp

import (
	"testing"
)

func TestEncodeString(t *testing.T) {
	fromUint := NewRLPStringFromUint(12111111111111111111)
	bytes := encodeString(fromUint.value)
	t.Log(bytes)
}