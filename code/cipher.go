package code

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

// 保持32位
func encodeCipherKey(key string) []byte {
	k, _ := hex.DecodeString(key)
	k32 := base64.StdEncoding.EncodeToString(k)
	k32 = Pad(k32, "-", 32)
	return []byte(k32[:32])
}

// 必须要用 hex.EncodeToString([]byte) 转为可读字符串
func EncryptCipher(text string, key string) (string, error) {
	// 进去怎么进去，出来就怎么出来就行了。如果是用 hex.DecodeString() 转换的，decrypt取出来的时候，就需要用 hex.EncodeToString()
	plaintext := []byte(text)
	k32 := encodeCipherKey(key)
	c, err := aes.NewCipher(k32)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	seal := gcm.Seal(nonce, nonce, plaintext, nil)
	// 额外增加的一步安全保障，防止 key 泄露后引起的问题
	seal[5], seal[len(seal)-1] = seal[len(seal)-1], seal[5]
	return base64.RawURLEncoding.EncodeToString(seal), nil
}

func DecryptCipher(text string, key string) (string, error) {
	ciphertext, err := base64.RawURLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < 6 {
		return "", errors.New("invalid cipher text " + text)
	}
	// 额外增加的一步安全保障，防止 key 泄露后引起的问题
	ciphertext[len(ciphertext)-1], ciphertext[5] = ciphertext[5], ciphertext[len(ciphertext)-1]

	k32 := encodeCipherKey(key)
	c, err := aes.NewCipher(k32)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	// 由于encrypt选择的是 []byte 转换，那么相应的，这里用 string() 转换回来
	return string(plaintext), err
}
