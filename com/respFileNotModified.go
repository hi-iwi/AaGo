package com

import (
	"net/http"
	"net/textproto"
	"strings"
	"time"
)

var unixEpochTime = time.Unix(0, 0)

// isZeroTime reports whether t is obviously unspecified (either zero or Unix()=0).
func isZeroTime(t time.Time) bool {
	return t.IsZero() || t.Equal(unixEpochTime)
}

/*
ETag   ==  Entity Tag
entity-tag = [ weak ] opaque-tag
weak      = %x57.2F ; "W/", case-sensitive
e.g.
	ETag: "abc"
	ETag: W/"abc"
	ETag: ""



First Time:
	Client   ---GET/HEAD /static/iwi.md  -->                                       [Server]
		<-- resp header Etag=2e681a-6-5d044840 ---

Next Time:
	[Client]   --- GET/HEAD /static/iwi.md,  header If-None-Match=2e681a-6-5d044840 -->     [Server]
			<-- if sever Etag == client If-None-Match, set If-None-Match=false, response 304
				else  set If-None-Match=true, response 200,




*/

// scanETag determines if a syntactically valid ETag is present at s. If so,
// the ETag and remaining text after consuming ETag is returned. Otherwise,
// it returns "", "".
func scanETag(s string) (etag string, remain string) {
	s = textproto.TrimString(s)
	start := 0
	if strings.HasPrefix(s, "W/") {
		start = 2
	}
	if len(s[start:]) < 2 || s[start] != '"' {
		return "", ""
	}
	// ETag is either W/"text" or "text".
	// See RFC 7232 2.3.
	for i := start + 1; i < len(s); i++ {
		c := s[i]
		switch {
		// Character values allowed in ETags.
		case c == 0x21 || c >= 0x23 && c <= 0x7E || c >= 0x80:
		case c == '"':
			return string(s[:i+1]), s[i+1:]
		default:
			return "", ""
		}
	}
	return "", ""
}

// etagStrongMatch reports whether a and b match using strong ETag comparison.
// Assumes a and b are valid ETags.
func etagStrongMatch(a, b string) bool {
	return a == b && a != "" && a[0] == '"'
}

// etagWeakMatch reports whether a and b match using weak ETag comparison.
// Assumes a and b are valid ETags.
func etagWeakMatch(a, b string) bool {
	return strings.TrimPrefix(a, "W/") == strings.TrimPrefix(b, "W/")
}

// condResult is the result of an HTTP request precondition check.
// See https://tools.ietf.org/html/rfc7232 section 3.
type condResult int

const (
	condNone condResult = iota
	condTrue
	condFalse
)

func (resp *RespStruct) checkIfMatch() condResult {
	imp, e := resp.req.Header("If-Match")
	if e != nil || !imp.NotEmpty() {
		return condNone
	}
	im := imp.String()
	for {
		im = textproto.TrimString(im)
		if len(im) == 0 {
			break
		}
		if im[0] == ',' {
			im = im[1:]
			continue
		}
		if im[0] == '*' {
			return condTrue
		}
		etag, remain := scanETag(im)
		if etag == "" {
			break
		}
		etg, e := resp.req.Header("Etag")
		if e == nil && etagStrongMatch(etag, etg.String()) {
			return condTrue
		}
		im = remain
	}

	return condFalse
}

func (resp *RespStruct) checkIfUnmodifiedSince(modtime time.Time) condResult {
	iusp, e := resp.req.Header("If-Unmodified-Since")
	if e != nil || !iusp.NotEmpty() || isZeroTime(modtime) {
		return condNone
	}
	ius := iusp.String()
	if t, err := http.ParseTime(ius); err == nil {
		// The Date-Modified header truncates sub-second precision, so
		// use mtime < t+1s instead of mtime <= t to check for unmodified.
		if modtime.Before(t.Add(1 * time.Second)) {
			return condTrue
		}
		return condFalse
	}
	return condNone
}

func (resp *RespStruct) checkIfNoneMatch() condResult {
	inmp, e := resp.req.Header("If-None-Match")

	if e != nil || !inmp.NotEmpty() {
		return condNone
	}
	buf := inmp.String()
	for {
		buf = textproto.TrimString(buf)
		if len(buf) == 0 {
			break
		}
		if buf[0] == ',' {
			buf = buf[1:]
		}
		if buf[0] == '*' {
			return condFalse
		}
		etag, remain := scanETag(buf)
		if etag == "" {
			break
		}
		etg, e := resp.req.Header("Etag")
		if e == nil && etagWeakMatch(etag, etg.String()) {
			return condFalse
		}
		buf = remain
	}
	return condTrue
}

func (resp *RespStruct) checkIfModifiedSince(modtime time.Time) condResult {

	if resp.req.Method != "GET" && resp.req.Method != "HEAD" {
		return condNone
	}
	imsp, e := resp.req.Header("If-Modified-Since")

	if e != nil || !imsp.NotEmpty() || isZeroTime(modtime) {
		return condNone
	}
	t, err := http.ParseTime(imsp.String())
	if err != nil {
		return condNone
	}
	// The Date-Modified header truncates sub-second precision, so
	// use mtime < t+1s instead of mtime <= t to check for unmodified.
	if modtime.Before(t.Add(1 * time.Second)) {
		return condFalse
	}
	return condTrue
}

func (resp *RespStruct) checkIfRange(modtime time.Time) condResult {
	if resp.req.Method != "GET" {
		return condNone
	}
	irp, e := resp.req.Header("If-Range")
	if e != nil || !irp.NotEmpty() {
		return condNone
	}
	etag, _ := scanETag(irp.String())
	if etag != "" {
		etg, e := resp.req.Header("Etag")
		if e == nil && etagStrongMatch(etag, etg.String()) {
			return condTrue
		}
		return condFalse
	}
	// The If-Range value is typically the ETag value, but it may also be
	// the modtime date. See golang.org/issue/8367.
	if modtime.IsZero() {
		return condFalse
	}
	t, err := http.ParseTime(irp.String())
	if err != nil {
		return condFalse
	}
	if t.Unix() == modtime.Unix() {
		return condTrue
	}
	return condFalse
}

// checkPreconditions evaluates request preconditions and reports whether a precondition
// resulted in sending StatusNotModified or StatusPreconditionFailed.
func (resp *RespStruct) checkPreconditions(modtime time.Time) (done bool, rangeHeader string) {
	// This function carefully follows RFC 7232 section 6.
	r := resp.req
	ch := resp.checkIfMatch()
	if ch == condNone {
		ch = resp.checkIfUnmodifiedSince(modtime)
	}
	if ch == condFalse {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return true, ""
	}
	switch resp.checkIfNoneMatch() {
	case condFalse:
		if r.Method == "GET" || r.Method == "HEAD" {
			resp.code = 403
			return true, ""
		}
		resp.code = 412
		return true, ""
	case condNone:
		if resp.checkIfModifiedSince(modtime) == condFalse {
			resp.code = 403
			return true, ""
		}
	}
	rhp, e := r.Header("Range")
	if e != nil && !rhp.NotEmpty() {
		rangeHeader = rhp.String()
		if resp.checkIfRange(modtime) == condFalse {
			rangeHeader = ""
		}
	}
	return false, rangeHeader
}
