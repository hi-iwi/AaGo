package aenum

type Status3 int8

const (
	Fail Status3 = -1 // 失败、拒绝
	Pend Status3 = 0  // 待定
	Ok   Status3 = 1  // 成功、同意
)
