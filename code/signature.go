package code

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"sort"
)

// SortedParams
func SortedParams(params map[string]string, ignoredSignKey string, joint bool) string {
	signStr := ""
	keys := make([]string, 0, len(params)-1)
	for k, _ := range params {
		if k != ignoredSignKey {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for i, k := range keys {
		if params[k] != "" {
			if joint {
				// 用 = & 连接
				if i > 0 {
					signStr += "&"
				}
				signStr += k + "=" + params[k]
			} else {
				signStr += k + params[k]
			}
		}
	}

	return signStr
}

// WriteSortedParams
func WriteSortedParams(w *bufio.Writer, params map[string]string, ignoredSignKey string, joint bool) {
	keys := make([]string, 0, len(params)-1)
	for k, _ := range params {
		if k != ignoredSignKey {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for i, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		if joint {
			// 用 = & 连接
			if i > 0 {
				w.WriteByte('&')
			}
			w.WriteString(k)
			w.WriteByte('=')
			w.WriteString(v)
		} else {
			w.WriteString(k)
			w.WriteString(v)
		}
	}
	return
}

// Sha1Signature Sha1 签名，大写结果
func Sha1Signature(params map[string]string, ignoredSignKey string, bufsize int, joint bool) string {
	h1 := sha1.New()
	if bufsize > 0 {
		// specify memory size
		bufw := bufio.NewWriterSize(h1, bufsize)
		WriteSortedParams(bufw, params, ignoredSignKey, joint)
		bufw.Flush()
	} else {
		// not specify memory size
		p := SortedParams(params, ignoredSignKey, joint)
		io.WriteString(h1, p)
	}
	bs := make([]byte, hex.EncodedLen(h1.Size()))
	hex.Encode(bs, h1.Sum(nil))
	return string(bytes.ToUpper(bs))
}

// HmacSignature HMAC 签名
// e.g. HmacSignature(sha1.New(), )
//func HmacSignature(h func() hash.Hash, params map[string]interface{}) string {
//	p := []byte("s")
//	mac := hmac.New(h, p)
//	mac.Write()
//}
