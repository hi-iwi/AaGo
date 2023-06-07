package afmt

import (
	"html/template"
	"sort"
	"strings"
)

func Url(url string, params map[string]string) template.URL {
	var s strings.Builder
	s.WriteString(url)
	n := strings.IndexByte(url, '?')
	if n > 0 {
		g := url[len(url)-1]
		if g != '&' && g != '?' {
			s.WriteByte('&')
		}
	} else {
		s.WriteByte('?')
	}
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v != "" && v != "0" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	for i, k := range keys {
		s.WriteString(k)
		s.WriteByte('=')
		s.WriteString(params[k])
		if i < len(keys)-1 {
			s.WriteByte('&')
		}
	}
	return template.URL(s.String())
}
