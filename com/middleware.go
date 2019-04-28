package com

import (
	"github.com/kataras/iris"
	"github.com/luexu/randm"
)

func Middleware(ictx iris.Context) {
	defer ictx.Next()

	ictx.Values().Set("traceid", randm.NewUUID().String())
}
