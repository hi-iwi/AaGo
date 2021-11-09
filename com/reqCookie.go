package com

import "net/http"

func (r *Req) Cookie(name string) (*http.Cookie, error) {
	return r.r.Cookie(name)
}

func (r *Req) AddCookie(c *http.Cookie) {
	r.r.AddCookie(c)
}

func (r *Req) Cookies() []*http.Cookie {
	return r.r.Cookies()
}
