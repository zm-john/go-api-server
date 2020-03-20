package utils

import "testing"

func Test_Token(t *testing.T) {
	sub := "2"
	conf := JWTConfig{Key: "zxcvfd", TTL: 3600}
	token, err := NewToken(sub, conf)

	if err != nil {
		t.Error(err)
	}

	_sub, err := ParseToken(token.Value, conf)
	if err != nil {
		t.Error(err)
	}

	if sub != _sub {
		t.Fail()
	}
}
