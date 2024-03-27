package com

type RespContentDTO struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Payload any    `json:"data"`
}
