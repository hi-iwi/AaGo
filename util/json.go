package util

import (
	"bytes"
	"encoding/json"
)

// 统一json
// @TODO resp 里面的json，不能ignore  json:"-"

func JsonString(v interface{}) ([]byte, error) {
	// json Marshal 不转译 HTML 字符
	buf := bytes.NewBuffer([]byte{})
	je := json.NewEncoder(buf)
	je.SetEscapeHTML(false)

	if err := je.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
