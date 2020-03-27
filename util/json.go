package util

import (
	"bytes"
	"encoding/json"
)

// 统一json
// @TODO resp 里面的json，不能ignore  json:"-"  和 omerti ； 解析 bool 为 字符串
// json.Marshal 默认 escapeHtml 为true,会转义 <、>、&
func JsonString(v interface{}) ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})
	je := json.NewEncoder(buf)
	je.SetEscapeHTML(false) // json Marshal 不转译 HTML 字符

	if err := je.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
