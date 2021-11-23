package com_test

import (
	"encoding/json"
	"github.com/hi-iwi/AaGo/util"
	"testing"
	"time"
)

type xJson struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
	KO   float64 `json:"ko"`
}
type yJson struct {
	xJson
}
type subJson struct {
	yJson
}

type innerJson struct {
	subJson
	Name string `json:"name"`
}
type mainJson struct {
	innerJson
	Html string `json:"html"`
	Age  int    `json:"main_age"`
	Test string `json:"-"`
}

func TestJson(t *testing.T) {

	xj := xJson{"iwi", 18, 1.34}
	yj := yJson{xJson: xj}
	sj := subJson{yJson: yj}
	ij := innerJson{subJson: sj, Name: "Iwi"}
	html := `
		<html>
			<body class="hello">${.World}</body>
		</html>
`
	var (
		tm int64
		ms []byte
		x  int64
	)
	mj := mainJson{innerJson: ij, Html: html, Test: "LOVESS"}

	tm = time.Now().UnixNano()
	ms, _ = util.JsonString(mj)
	x = time.Now().UnixNano() - tm
	t.Logf("%d util.JsonString:\n %s", x, string(ms))

	tm = time.Now().UnixNano()
	ms, _ = json.Marshal(mj)
	x = time.Now().UnixNano() - tm
	t.Logf("%d json.Marshal:\n %s", x, string(ms))

	tm = time.Now().UnixNano()
	ms, _ = util.JsonString(mj)
	x = time.Now().UnixNano() - tm
	t.Logf("%d util.JsonString:\n %s", x, string(ms))

	tm = time.Now().UnixNano()
	ms, _ = json.Marshal(mj)
	x = time.Now().UnixNano() - tm
	t.Logf("%d json.Marshal:\n %s", x, string(ms))

}
