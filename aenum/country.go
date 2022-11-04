package aenum

// 基础类，后面可以继承扩展
type Country struct {
	Id          CountryId `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`         // 国家码，通用 ISO 3166-1 三位/二位字母代码；若采用二位，跟域名后缀一致
	CallingCode string    `json:"calling_code"` // 国际电话码
	Icon        string    `json:"icon"`         // image url  or icon font
}
