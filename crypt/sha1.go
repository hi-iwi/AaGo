package crypt

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func Sha1(str string, length ...int) string {
	h := sha1.New()
	io.WriteString(h, str)
	r := hex.EncodeToString(h.Sum(nil))
	if len(length) > 0 && length[0] > 0 {
		if length[0] > len(r) {
			r = Pad(r, length[0])
		} else {
			r = r[0:length[0]]
		}

	}
	return r
}
