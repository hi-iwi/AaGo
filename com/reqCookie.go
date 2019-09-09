package com

import "net/http"

func (req *Req) Cookie(name string) (*http.Cookie, error) {
	return req.r.Cookie(name)
}

func (req *Req) AddCookie(c *http.Cookie) {
	req.r.AddCookie(c)
}

func (req *Req) Cookies() []*http.Cookie {
	return req.r.Cookies()
}
