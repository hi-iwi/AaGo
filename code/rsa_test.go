package code_test

import (
	"testing"
)

// 配合jsencrypt.js使用
func TestJsEncryptRSA(t *testing.T) {
	privk, pubk, err := GenRSA(2048)
	privkey := string(privk)
	pubkey := string(pubk)

	msg := []byte("Luexu.com")
	en, err := RsaEncryptPKCS1(msg, []byte(pubkey))
	if err != nil {
		t.Error(err)
	}

	_, err = RsaDecryptPKCS1(en, []byte(privkey))
	if err != nil {
		t.Error(err)
	}
}
