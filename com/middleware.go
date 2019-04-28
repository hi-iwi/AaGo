package com

import (
	"github.com/kataras/iris"
	"github.com/luexu/randm"
)

func Middleware(ictx iris.Context) {
	ictx.Values().Set("traceid", randm.NewUUID().String())
}
