package com

import "github.com/kataras/iris/v12"

func QueryToViewData(ictx iris.Context, name string) {
	r := NewReq(ictx)
	v, _ := r.QueryString(name, false)
	ictx.ViewData(name, v)
}
