package code

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// JSEncrypt RSA 使用的是标准的PKCS1，所以对其加密、解密需要用这个

/*
bits 2048
privkey, err := rsa.GenerateKey(rand.Reader, bits)   生成privkey, 而 pubkey = privkey.PublicKey

*/

// GenRSA  如 openssl genrsa -out rsa_2048_priv.pem 2048
// 生成的是已经base64后的字节，直接用 string()转化即可
func GenRSA(bits int) ([]byte, []byte, error) {
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	privkey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	pub := &priv.PublicKey
	pkix, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, nil, err
	}
	pubkey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pkix,
	})

	return privkey, pubkey, nil
}

func RsaBytesToPrivateKey(privkey []byte) (priv *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(privkey)
	b := block.Bytes
	if enc := x509.IsEncryptedPEMBlock(block); enc {
		if b, err = x509.DecryptPEMBlock(block, nil); err != nil {
			return nil, err
		}
	}
	return x509.ParsePKCS1PrivateKey(b)
}

// BytesToPublicKey bytes to public key
func RsaBytesToPublicKey(pubkey []byte) (pub *rsa.PublicKey, err error) {
	block, _ := pem.Decode(pubkey)
	b := block.Bytes
	if enc := x509.IsEncryptedPEMBlock(block); enc {
		if b, err = x509.DecryptPEMBlock(block, nil); err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}

	if pub, ok := ifc.(*rsa.PublicKey); ok {
		return pub, nil
	}
	return nil, errors.New("bytes to public key error")
}

// OAEP 加密
func RsaEncryptOAEP(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	return rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
}

func RsaDecryptOAEP(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	return rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
}

// PKCS1   适用于 JSEncrypt
// 这个生成的是二进制字节，可以用base64将结果转化下返回给客户端；注意保证客户端跟服务端base64解码方式一样
func RsaEncryptPKCS1(msg []byte, pubkey []byte) ([]byte, error) {
	pub, err := RsaBytesToPublicKey(pubkey)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
}

func RsaDecryptPKCS1(ciphertext []byte, privkey []byte) ([]byte, error) {
	priv, err := RsaBytesToPrivateKey(privkey)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
