package aenum

// 中国民族
type Ethnicity uint8

const (
	UnknownNationality Ethnicity = 0
	HanZu              Ethnicity = 1  // 汉族
	ManZu              Ethnicity = 2  // 满族
	MengguZu           Ethnicity = 3  // 蒙古族
	HuiZu              Ethnicity = 4  // 回族
	ZangZu             Ethnicity = 5  // 藏族
	WeiwuerZu          Ethnicity = 6  // 维吾尔族
	MiaoZu             Ethnicity = 7  // 苗族
	YiZu               Ethnicity = 8  // 彝族
	ZhuangZu           Ethnicity = 9  // 壮族
	BuyiZu             Ethnicity = 10 // 布依族

)
