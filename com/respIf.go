package com

type RespContentDTO struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"data"`
}

// @TODO
// ?_map=time,service,connections:[name,scheme],server_id,test:{a,b,c}
func (resp RespStruct) handlePayload(a interface{}) interface{} {
	return a
}
