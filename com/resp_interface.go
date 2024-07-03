package com

type RespContentDTO struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Payload any    `json:"data"`
}

func OK(payload any) RespContentDTO {
	return RespContentDTO{
		Code:    200,
		Msg:     "OK",
		Payload: payload,
	}
}
