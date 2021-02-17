package code_test

import (
	"github.com/hi-iwi/AaGo/code"
	"testing"
)

// 配合jsencrypt.js使用
func TestJsEncryptRSA(t *testing.T) {
	privk, pubk, err := code.GenRSA(2048)
	privkey := string(privk)
	pubkey := string(pubk)

	msg := []byte("Luexu.com")
	en, err := code.RsaEncryptPKCS1(msg, []byte(pubkey))
	if err != nil {
		t.Error(err)
	}

	_, err = code.RsaDecryptPKCS1(en, []byte(privkey))
	if err != nil {
		t.Error(err)
	}
}
